# ansible-dappley-local
 Testing go-dappley locally (127.0.0.1) on an aws EC2 instance

### Local Testing Pipeline:
```
pipeline {
    agent any
    tools {
        go 'go-1.16.6'
    }
    environment {
        GO1166MODULE = 'on'
    }
    parameters {
        gitParameter branchFilter: 'origin/(.*)', defaultValue: 'master',  name: 'MASTER',  type: 'PT_BRANCH'
        gitParameter branchFilter: 'origin/(.*)', defaultValue: 'develop', name: 'DEVELOP', type: 'PT_BRANCH'
    }
    stages {
        stage('SCM Checkout Master Branch') {
            steps {
                git branch: "${params.MASTER}", url: 'https://github.com/dappley/go-dappley.git'
            }
        }
        stage('Make Master Branch') {
            steps {
                sh 'make build'
            }
        }
        stage('Test Master Branch') {
            steps {
                sh 'mkdir master'
                sh 'git show > master/change.txt'
                sh 'make testall > master/log.txt'
                sh 'git log --pretty=fuller HEAD^..HEAD > master/commitInfo.txt'
            }
        }
        stage('Move Master Files') {
            steps {
                sh 'mv master ../send-test-results'
            }
        }
        stage('Clear Master Directory') {
            steps {
                sh 'rm -r *'
            }
        }
        stage('SCM Checkout Develop Branch') {
            steps {
                git branch: "${params.DEVELOP}", url: 'https://github.com/dappley/go-dappley.git'
            }
        }
        stage('Compile Develop Branch') {
            steps {
                sh 'make build'
            }
        }
        stage('Test Develop Branch') {
            steps {
                sh 'mkdir develop'
                sh 'git show > develop/change.txt'
                sh 'make testall > develop/log.txt'
                sh 'git log --pretty=fuller HEAD^..HEAD > develop/commitInfo.txt'
            }
        }
        stage('Move Develop Files') {
            steps {
                sh 'mv develop ../send-test-results'
            }
        }
        stage('Clear Develop Directory') {
            steps {
                sh 'rm -r *'
            }
        }
    }
}
```

### Single Server Testing Pipeline:
```
pipeline {
    agent any
    tools {
        go 'go-1.16.6'
    }
    environment {
        GO1166MODULE = 'on'
    }
    stages {
        stage('SCM Checkout') {
            steps {
                git 'https://github.com/heesooh/go-dappley-single-server-tester/'
            }
        }
        stage('Create Nodes') {
            steps {
                sh "aws ec2 run-instances --image-id ami-09e67e426f25ce0d7 --instance-type m5.large --count 1 --key-name jenkins --security-group-ids sg-0805d37807366482d --tag-specifications 'ResourceType=instance,Tags=[{Key=Name,Value=JenkinsLocalTestServer}]' > local_test_server.txt"
            }
        }
        stage('Build') {
            steps {
                sh 'go mod init github.com/heesooh/go-dappley-single-server-tester'
                sh 'go mod tidy'
                sh 'go build'
            }
        }
        stage('Create Directory') {
            steps {
                sh 'mkdir ../send-test-results/master'
                sh 'mkdir ../send-test-results/develop'
            }
        }
        stage('Update Hosts') {
            steps {
                sh './go-dappley-single-server-tester -function update'
            }
        }
        stage('Wait 10 Seconds') {
            steps {
                sh 'sleep 10'
            }
        }
        stage('Initialize Hosts') {
            steps {
                sh './go-dappley-single-server-tester -function initialize'
            }
        }
        stage('Setup Host Nodes') {
            steps {
                sh 'ansible-playbook ./playbooks/setup.yml'
            }
        }
        stage('Test Master & Develop') {
            steps {
                sh 'ansible-playbook ./playbooks/test.yml'
            }
        }
        stage('Terminate Host Nodes') {
            steps {
                sh './go-dappley-single-server-tester -function terminate'
            }
        }
        stage('Close') {
            steps {
                sh 'rm -r *'
            }
        }
    }
}
```