package commands

import (
	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
)

type CommandExecFunc func(meta common.CommandMetadata) (uint16, string)

type ServerCommand struct {
	Name            string
	ExecuteFunction CommandExecFunc
}

func CreateServerCommand(name string, exec CommandExecFunc) ServerCommand {
	return ServerCommand{
		name,
		exec,
	}
}

func (c ServerCommand) Execute(sender *common.Connection, args []byte) (uint16, string) {
	return c.ExecuteFunction(common.NewCommandMetadata(sender, args))
}

func (c ServerCommand) GetName() string {
	return c.Name
}
