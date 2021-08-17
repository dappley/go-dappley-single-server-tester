package main 

import (
	"fmt"
	"flag"
)

func main() {
	var function, senderEmail, senderPasswd string
	flag.StringVar(&function, "function", "default_function", "Name of the function that will be run.")
	flag.StringVar(&senderEmail,  "senderEmail",  "default_email@example.com", "Email of the addressee.")
	flag.StringVar(&senderPasswd, "senderPasswd", "default_password", "Email password of the addressee.")
	flag.Parse()

	if function == "update" {
		update_host()
	} else if function == "initialize" {
		initialize_host()
	} else if function == "ssh_command" {
		ssh_command()
	} else if function == "terminate" {
		terminate_host()
	} else {
		fmt.Println("Function Invalid!\nFunction can be one of below:\nupdate\ninitialize\nssh_command\nterminate")
	}
}