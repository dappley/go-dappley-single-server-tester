package aws 

import (
	"github.com/heesooh/go-dappley-single-server-testing/helper"
	"io/ioutil"
	"strings"
	"bufio"
	"log"
	"os"
)

//Adds the server information to the hosts and instance_ids file.
func Update_host() {
	var private_ips, instance_ids string
	fileName := "local_test_server.txt"

	//Create txt files for server info.
	host_file, err := os.Create("hosts")
	if err != nil { log.Fatal("Unable to create the \"hosts\" file!") }
	id_file, err := os.Create("instance_ids")
	if err != nil { log.Fatal("Unbale to create the \"instance_ids\" file!") }

	//Read local_test_server.txt
	node_byte, err := ioutil.ReadFile(fileName)
	if err != nil { log.Fatal("Failed to read", fileName, "!") }

	scanner := bufio.NewScanner(strings.NewReader(string(node_byte)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "InstanceId") {
			instance_ids += helper.TrimLeftRight(line, "\"", "\",") + "\n"
		}
		if strings.Contains(line, "PrivateIpAddress") {
			private_ips += "[LOCAL_TEST_SERVER]\n" + helper.TrimLeftRight(line, "\"", "\",") + "\n"
			break
		}
	}

	//Write server info to txt files.
	_, err = host_file.WriteString(private_ips)
	if err != nil { log.Fatal("Unable to write on file \"hosts\"!") }
	_, err = id_file.WriteString(instance_ids)
	if err != nil { log.Fatal("Unable to write on file \"instance_ids\"!") }
}