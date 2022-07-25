package server

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
	"github.com/HappyGick/filebroadcastprotocol/protocol/common/util"
	"github.com/HappyGick/filebroadcastprotocol/protocol/server/commands"
)

type CommandSuccessFunc func()
type CommandFailFunc func(err string)

type ServerCommandHandler struct {
	commands map[string]commands.ServerCommand
}

func NewServerCommandHandler() ServerCommandHandler {
	return ServerCommandHandler{
		map[string]commands.ServerCommand{},
	}
}

func (ch *ServerCommandHandler) RegisterCommand(newCommand commands.ServerCommand) {
	ch.commands[newCommand.GetName()] = newCommand
}

func (ch *ServerCommandHandler) Handle(command []byte, sender *common.Connection) error {
	name, args, err := preprocessCommand(command)

	if err != nil {
		return err
	}

	if val, ok := ch.commands[name]; ok {
		code, msg := val.Execute(sender, args)
		fmt.Println("Command result:", code, msg)
	} else {
		sender.Send([]byte("invalid command"))
		fmt.Println("Received invalid command:", name)
	}

	return nil
}

func preprocessCommand(command []byte) (string, []byte, error) {
	var name bytes.Buffer
	var args bytes.Buffer
	r := bytes.NewReader(command)
	w := bufio.NewWriter(&name)
	buf, err := util.ReadUntil(r, 0)
	if err != nil {
		return "", nil, err
	}
	w.Write(buf)
	w.Flush()
	w = bufio.NewWriter(&args)
	r.WriteTo(w)
	w.Flush()

	return name.String(), args.Bytes(), nil
}
