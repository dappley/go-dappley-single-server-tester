package main 

import (
	"log"
	"fmt"
	"flag"
	"errors"
)

func main() {
	var function, senderEmail, senderPasswd string
	flag.StringVar(&function, "function", "default_function", "Name of the function that will be run.")
	flag.StringVar(&senderEmail,  "senderEmail",  "default_email", "Email of the addressee.")
	flag.StringVar(&senderPasswd, "senderPasswd", "default_password", "Email password of the addressee.")
	flag.Parse()

	err := checkFlags(function, senderEmail, senderPasswd)
	if err != nil {
		log.Fatal(err)
		return
	}

	if function == "update" {
		update_host()
	} else if function == "initialize" {
		initialize_host()
	} else if function == "ssh_command" {
		ssh_command()
	} else if function == "terminate" {
		terminate_host()
	} else {
		fmt.Println("Error: Function is invalid!\nFunction can be one of below:\nupdate\ninitialize\nssh_command\nterminate")
	}
}

//----------Helper----------
func checkFlags(function string, email string, password string) (err error) {
	switch {
	case function == "default_function":
		err = errors.New("Error: Function is missing!")
	case email == "default_email":
		err = errors.New("Error: Email is missing!")
	case password == "default_password":
		err = errors.New("Error: Password is missing!")
	default:
		err = nil
	}
	return err
}