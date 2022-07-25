package common

import (
	"time"
)

type CommandMetadata struct {
	sender       *Connection
	args         []byte
	timeReceived time.Time
}

func NewCommandMetadata(sender *Connection, args []byte) CommandMetadata {
	return CommandMetadata{
		sender,
		args,
		time.Now(),
	}
}

func (c CommandMetadata) GetSender() *Connection {
	return c.sender
}

func (c CommandMetadata) GetArgs() []byte {
	return c.args
}
