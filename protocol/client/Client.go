package client

import (
	"fmt"
	"net"

	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
	"github.com/HappyGick/filebroadcastprotocol/protocol/common/util"
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

func (c *FileBroadcastClient) JoinChannel(channel uint64) ([]byte, error) {
	msg := util.NewByteBuffer([]byte{}).Append([]byte("CHAN\x00")).AppendUint64BE(channel)
	err := util.ReportError(msg)

	if err != nil {
		return nil, err
	}

	err = c.send(msg.Bytes())

	if err != nil {
		return nil, err
	}

	return c.receive()
}

func (c *FileBroadcastClient) SendFile(channel uint64, path string) ([]byte, error) {
	file, err := common.FileRef(path)

	if err != nil {
		return nil, err
	}

	data := util.NewByteBuffer([]byte{}).AppendFile(file)
	name := file.GetName()
	sizebytes := util.NewByteBuffer([]byte{}).AppendUint64BE(file.GetSize())
	channelbytes := util.NewByteBuffer([]byte{}).AppendUint64BE(channel)
	msg := util.NewByteBuffer([]byte{}).AppendMultiple(
		[]byte("SEND\x00"),
		channelbytes.Bytes(),
		[]byte(name+"\x00"),
		sizebytes.Bytes(),
		data.Bytes(),
	)

	err = util.ReportError(data, sizebytes, channelbytes, msg)

	if err != nil {
		return nil, err
	}

	err = c.send(msg.Bytes())

	if err != nil {
		return nil, err
	}

	return c.receive()
}
