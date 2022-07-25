package common

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Connection struct {
	conn     *net.Conn
	lastPing time.Time
}

func NewConnection(conn *net.Conn) Connection {
	return Connection{
		conn,
		time.Now(),
	}
}

func (c *Connection) Receive() ([]byte, error) {
	c.lastPing = time.Now()
	buflen := make([]byte, 8)
	_, err := (*c.conn).Read(buflen)
	if err != nil {
		return []byte{}, err
	}
	buf := make([]byte, binary.BigEndian.Uint64(buflen))
	len, err := (*c.conn).Read(buf)
	fmt.Println("Received", len, "bytes from", c.GetAddr())
	return buf, err
}

func (c Connection) Send(data []byte) error {
	length := make([]byte, 8)
	binary.BigEndian.PutUint64(length, uint64(len(data)))
	len, err := (*c.conn).Write(append(length, data...))
	fmt.Println("Sent", len, "bytes to", c.GetAddr())
	return err
}

func (c Connection) Close() error {
	return (*c.conn).Close()
}

func (c Connection) GetAddr() string {
	return (*c.conn).RemoteAddr().String()
}
