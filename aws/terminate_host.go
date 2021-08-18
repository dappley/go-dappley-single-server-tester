package aws 

import (
	"github.com/heesooh/go-dappley-single-server-testing/helper"
	"io/ioutil"
	"strings"
	"bufio"
	"fmt"
	"log"
)

//Termiante all servers via aws cli command.
func Terminate_host() {
	fileName := "instance_ids"
	instance_byte, err := ioutil.ReadFile(fileName)
	if err != nil { log.Fatal("Failed to read", fileName, "!") }

	scanner := bufio.NewScanner(strings.NewReader(string(instance_byte)))
	for scanner.Scan() {
		terminate_instance := "aws ec2 terminate-instances --instance-ids " + scanner.Text()
		output := helper.ShellCommandExecuter(terminate_instance)
		fmt.Printf("%s\n", output)
		fmt.Println(terminate_instance)
	}
}