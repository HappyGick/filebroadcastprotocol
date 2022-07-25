package main

import (
	"fmt"

	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
	"github.com/HappyGick/filebroadcastprotocol/protocol/server"
)

func handleConnection(conn common.Connection) {
	defer conn.Close()
	data, err := conn.Receive()
	if err != nil {
		fmt.Println("Error receiving data:", err)
		return
	}
	fmt.Println("Received:", data)
	conn.Send([]byte("success"))
}

func main() {
	fmt.Println("Starting")
	server := server.Create(server.DefaultOptions())
	server.Listen()
}
