package server

import (
	"bytes"
	"fmt"

	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
	"github.com/HappyGick/filebroadcastprotocol/protocol/common/util"
	"github.com/HappyGick/filebroadcastprotocol/protocol/server/commands"
)

type ServerCommandHandler struct {
	commands map[string]commands.CommandBase
}

func NewServerCommandHandler() ServerCommandHandler {
	return ServerCommandHandler{
		map[string]commands.CommandBase{},
	}
}

func RegisterCommand[T any](ch *ServerCommandHandler, newCommand commands.ServerCommand[T]) {
	ch.commands[newCommand.GetName()] = newCommand
}

func (ch *ServerCommandHandler) Handle(command []byte, sender *common.Connection) error {
	name, args, err := preprocessCommand(command)

	if err != nil {
		return err
	}

	if val, ok := ch.commands[name]; ok {
		// por qué no tiene polimorfismo paramétrico? wtf
		code, msg := val.Execute(sender, args)
		fmt.Println("Command result:", code, msg)
	} else {
		sender.Send([]byte("invalid command"))
		fmt.Println("Received invalid command:", name)
	}

	return nil
}

func preprocessCommand(command []byte) (string, []byte, error) {
	r := bytes.NewReader(command)
	name := util.NewByteBuffer([]byte{}).AppendUntil(r, 0)
	args := util.NewByteBuffer([]byte{}).AppendAll(r)

	return name.String(), args.Bytes(), util.ReportError(name, args)
}
