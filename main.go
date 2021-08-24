package main 

import (
	"github.com/heesooh/go-dappley-single-server-testing/helper"
	"github.com/heesooh/go-dappley-single-server-testing/aws"
	"flag"
	"fmt"
	"log"
)

func main() {
	var function, senderEmail, senderPasswd string
	flag.StringVar(&function, "function", "default_function", "Name of the function that will be run.")
	flag.StringVar(&senderEmail,  "senderEmail",  "default_email", "Email of the addressee.")
	flag.StringVar(&senderPasswd, "senderPasswd", "default_password", "Email password of the addressee.")
	flag.Parse()

	err := helper.CheckFlags(senderEmail, senderPasswd)
	if err != nil {
		log.Fatal(err)
		return
	}

	if function == "update" {
		aws.Update_host()
	} else if function == "initialize" {
		aws.Initialize_host()
	} else if function == "ssh_command" {
		aws.SSH_command()
	} else if function == "terminate" {
		aws.Terminate_host()
	} else if function == "default_function" {
		log.Fatal("Error: Function is missing!")
	} else {
		fmt.Println("Error: Function is invalid!\nFunction can be one of below:\nupdate\ninitialize\nssh_command\nterminate")
	}
}