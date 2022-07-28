package common

import (
	"time"
)

type CommandSuccessFunc[T any] func(args T)
type CommandFailFunc func(err error)

type CommandMetadata[ReturnType any] struct {
	sender       *Connection
	args         []byte
	timeReceived time.Time
	onSuccess    CommandSuccessFunc[ReturnType]
	onFail       CommandFailFunc
}

func NewCommandMetadata[ReturnType any](sender *Connection, args []byte, successFunc CommandSuccessFunc[ReturnType], failFunc CommandFailFunc) CommandMetadata[ReturnType] {
	return CommandMetadata[ReturnType]{
		sender,
		args,
		time.Now(),
		successFunc,
		failFunc,
	}
}

func (c CommandMetadata[ReturnType]) GetSender() *Connection {
	return c.sender
}

func (c CommandMetadata[ReturnType]) GetArgs() []byte {
	return c.args
}

func (c CommandMetadata[ReturnType]) Success(args ReturnType) {
	c.onSuccess(args)
}

func (c CommandMetadata[ReturnType]) Fail(err error) {
	c.onFail(err)
}
