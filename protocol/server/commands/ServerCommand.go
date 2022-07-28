package commands

import (
	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
)

type CommandExecFunc[T any] func(meta common.CommandMetadata[T]) (uint16, string)

type CommandBase interface {
	GetName() string
	Execute(sender *common.Connection, args []byte) (uint16, string)
}

type ServerCommand[ReturnType any] struct {
	CommandBase
	Name            string
	ExecuteFunction CommandExecFunc[ReturnType]
	onSuccess       common.CommandSuccessFunc[ReturnType]
	onFail          common.CommandFailFunc
}

func CreateServerCommand[ReturnType any](name string, exec CommandExecFunc[ReturnType], success common.CommandSuccessFunc[ReturnType], fail common.CommandFailFunc) ServerCommand[ReturnType] {
	return ServerCommand[ReturnType]{
		Name:            name,
		ExecuteFunction: exec,
		onSuccess:       success,
		onFail:          fail,
	}
}

func (c ServerCommand[ReturnType]) Execute(sender *common.Connection, args []byte) (uint16, string) {
	return c.ExecuteFunction(common.NewCommandMetadata(sender, args, c.onSuccess, c.onFail))
}

func (c ServerCommand[ReturnType]) GetName() string {
	return c.Name
}
