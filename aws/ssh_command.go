package aws 

import (
	"github.com/heesooh/go-dappley-single-server-testing/helper"
	"io/ioutil"
	"strings"
	"bufio"
	"fmt"
	"log"
)

//Prints out the ssh command for all servers.
func SSH_command() {	
	fileName := "instance_ids"
	instance_byte, err := ioutil.ReadFile(fileName)
	if err != nil { log.Fatal("Failed to read", fileName, "!") }

	scanner := bufio.NewScanner(strings.NewReader(string(instance_byte)))
	for scanner.Scan() {
		describe_instance := "aws ec2 describe-instances --instance-ids " + scanner.Text()
		output := helper.ShellCommandExecuter(describe_instance)

		description_scanner := bufio.NewScanner(strings.NewReader(string(output)))
		for description_scanner.Scan() {
			line := description_scanner.Text()
			if strings.Contains(line, "\"PublicIpAddress\":") {
				fmt.Println("ssh -i jenkins.pem ubuntu@" + helper.TrimLeftRight(line, "\"", "\","))
				break
			}
		}
	}
}