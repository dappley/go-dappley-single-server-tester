# ansible-dappley-local
 Testing go-dappley locally (127.0.0.1) on an aws EC2 instance

Pipeline:
```
pipeline {
    agent any
    tools {
        go 'go-1.16.3'
    }
    environment {
        GO1163MODULE = 'on'
    }
    stages {
        stage('SCM Checkout') {
            steps {
                git 'https://github.com/heesooh/ansible-dappley-local/'
            }
        }
        stage('Create Nodes') {
            steps {
                sh "aws ec2 run-instances --image-id ami-02701bcdc5509e57b --instance-type m5.xlarge --count 1 --key-name jenkins --security-group-ids sg-03d39dd5dc5ddeef7 --tag-specifications 'ResourceType=instance,Tags=[{Key=Name,Value=JenkinsLocalTestServer}]' > local_test_server.txt"
            }
        }
        stage('Build') {
            steps {
                sh 'go mod init github.com/heesooh/ansible-dappley-local'
                sh 'go mod tidy'
                sh 'go build'
            }
        }
        stage('Create Directory') {
            steps {
                sh 'mkdir ../go-dappley-commit-report/master'
                sh 'mkdir ../go-dappley-commit-report/develop'
            }
        }
        stage('Update Hosts') {
            steps {
                sh './ansible-dappley-local -function update'
            }
        }
        stage('Wait 10 Seconds') {
            steps {
                sh 'sleep 10'
            }
        }
        stage('Initialize Hosts') {
            steps {
                sh './ansible-dappley-local -function initialize'
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
                sh './ansible-dappley-local -function terminate'
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