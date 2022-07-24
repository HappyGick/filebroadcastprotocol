package client

import (
	"net"

	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
)

type FileBroadcastClient struct {
	connectAddress string
	conn           common.Connection
}

func New(address string) FileBroadcastClient {
	return FileBroadcastClient{
		address,
		common.Connection{},
	}
}

func (c *FileBroadcastClient) Connect() error {
	conn, err := net.Dial("tcp", c.connectAddress)
	if err == nil {
		c.conn = common.NewConnection(&conn)
	}
	return err
}

func (c FileBroadcastClient) Send(data []byte) error {
	return c.conn.Send(data)
}

func (c FileBroadcastClient) Receive() ([]byte, error) {
	return c.conn.Receive()
}

func (c *FileBroadcastClient) Disconnect() error {
	return c.conn.Close()
}
