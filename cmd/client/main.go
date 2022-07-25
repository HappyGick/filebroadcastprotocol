package main

import (
	"fmt"

	"github.com/HappyGick/filebroadcastprotocol/protocol/client"
)

func main() {
	client := client.New("localhost:3000")
	client.Connect()
	resp, err := client.SendFile(0, "/mnt/c/Users/famil/Documents/dev/gofileprotocol/bin/client/msg.txt")
	if err != nil {
		fmt.Println("Error sending file:", err)
	}
	fmt.Println("Response from server:", string(resp))
	err = client.Disconnect()
	if err != nil {
		fmt.Println("Error disconnecting:", err)
	}
}
