package main

import (
	"fmt"

	"github.com/HappyGick/filebroadcastprotocol/protocol/client"
)

func main() {
	client := client.New("localhost:3000")
	client.Connect()
	client.Send([]byte("hello world"))
	data, _ := client.Receive()
	fmt.Println(string(data))
	client.Disconnect()
}
