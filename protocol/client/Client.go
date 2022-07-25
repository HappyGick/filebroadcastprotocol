package client

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
)

type FileBroadcastClient struct {
	connectAddress string
	conn           common.Connection
	commandHandler ClientCommandHandler
}

func New(address string) FileBroadcastClient {
	return FileBroadcastClient{
		address,
		common.Connection{},
		ClientCommandHandler{},
	}
}

func (c *FileBroadcastClient) Connect() {
	conn, err := net.Dial("tcp", c.connectAddress)
	if err == nil {
		c.conn = common.NewConnection(&conn)
	} else {
		fmt.Println("Error connecting to", c.connectAddress, ":", err)
	}
}

func (c FileBroadcastClient) send(data []byte) error {
	return c.conn.Send(data)
}

func (c FileBroadcastClient) receive() ([]byte, error) {
	return c.conn.Receive()
}

func (c *FileBroadcastClient) Disconnect() error {
	return c.conn.Close()
}

func (c *FileBroadcastClient) SendFile(channel int, path string) ([]byte, error) {
	file, err := common.FileRef(path)

	if err != nil {
		return nil, err
	}

	data, err := file.Read()

	if err != nil {
		return nil, err
	}

	name := file.GetName()
	sizebytes := make([]byte, 8)
	binary.BigEndian.PutUint64(sizebytes, file.GetSize())
	channelbytes := make([]byte, 8)
	binary.BigEndian.PutUint64(channelbytes, uint64(channel))

	var msg bytes.Buffer
	w := bufio.NewWriter(&msg)
	w.Write([]byte("SEND\x00"))
	w.Write(channelbytes)
	w.Write([]byte(name + "\x00"))
	w.Write(sizebytes)
	w.Write(data)
	w.Flush()

	err = c.send(msg.Bytes())
	if err != nil {
		return nil, err
	}

	resp, err := c.receive()

	return resp, err
}
