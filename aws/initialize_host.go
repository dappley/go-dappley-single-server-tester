package aws 

import (
	"github.com/heesooh/go-dappley-single-server-tester/helper"
	"io/ioutil"
	"strings"
	"errors"
	"bufio"
	"fmt"
	"log"
)

//Runs until all servers are initialized.
func Initialize_host() {
	fileName := "instance_ids"
	instance_byte, err := ioutil.ReadFile(fileName)
	if err != nil { log.Fatal("Failed to read", fileName, "!") }

	scanner := bufio.NewScanner(strings.NewReader(string(instance_byte)))
	for scanner.Scan() {
		instance_id := scanner.Text()
		initializing := true
		fmt.Println("Initializing " + instance_id + "...")
		for initializing {
			initialize_instance := "aws ec2 describe-instance-status --instance-ids " + instance_id
			output := helper.ShellCommandExecuter(initialize_instance)

			status_scanner := bufio.NewScanner(strings.NewReader(string(output)))			
			for status_scanner.Scan() {
				line := status_scanner.Text()

				if strings.Contains(line, "\"InstanceStatuses\":") {
					status := helper.TrimLeftRight(line, "\"", "\"")
					if status == "[]" {
						err := errors.New("Instance " + instance_id + "has been termianted!")
						panic(err)
					}
				}
				if strings.Contains(line, "\"Status\":") {
					status := helper.TrimLeftRight(line, "\"", "\"")
					if status == "passed" {
						initializing = false
						fmt.Println("Instance " + instance_id + " initialized!")
						break
					}
				}
			}
		}
	}
}