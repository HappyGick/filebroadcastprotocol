package JoinChannel

import (
	"bytes"
	"fmt"

	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
	"github.com/HappyGick/filebroadcastprotocol/protocol/common/util"
	"github.com/HappyGick/filebroadcastprotocol/protocol/server/commands"
)

type JoinChannelReturn struct {
	ChannelId uint64
	Sender    *common.Connection
}

func NewJoinChannelCommand(onSuccess common.CommandSuccessFunc[JoinChannelReturn], onFail common.CommandFailFunc) commands.ServerCommand[JoinChannelReturn] {
	return commands.CreateServerCommand("CHAN", executeJoinCommand, onSuccess, onFail)
}

func interpretArgs(args []byte) (uint64, error) {
	buf := util.NewByteBuffer([]byte{}).AppendFrom(bytes.NewReader(args), 8)
	return buf.BEUint(), buf.GetError()
}

func executeJoinCommand(meta common.CommandMetadata[JoinChannelReturn]) (uint16, string) {
	chid, err := interpretArgs(meta.GetArgs())

	if err != nil {
		meta.Fail(err)
		return 1, "[" + meta.GetSender().GetAddr() + "] Error executing CHAN: " + err.Error()
	}

	meta.Success(JoinChannelReturn{ChannelId: chid, Sender: meta.GetSender()})
	meta.GetSender().Send([]byte("Successfully joined channel " + fmt.Sprint(chid)))

	return 0, ""
}
