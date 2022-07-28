package SendToChannel

import (
	"bytes"
	"fmt"

	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
	"github.com/HappyGick/filebroadcastprotocol/protocol/common/util"
	"github.com/HappyGick/filebroadcastprotocol/protocol/server/commands"
)

func NewSendToChannelCommand() commands.ServerCommand[any] {
	return commands.CreateServerCommand("SEND", executeSendCommand, nil, nil)
}

func interpretArgs(args []byte) (uint64, string, []byte, error) {
	r := bytes.NewReader(args)
	channel := util.NewByteBuffer([]byte{}).AppendFrom(r, 8)
	filename := util.NewByteBuffer([]byte{}).AppendUntil(r, 0)
	size := util.NewByteBuffer([]byte{}).AppendFrom(r, 8)
	data := util.NewByteBuffer([]byte{}).AppendFrom(r, size.BEUint())

	return channel.BEUint(), filename.String(), data.Bytes(), util.ReportError(channel, filename, size, data)
}

func executeSendCommand(meta common.CommandMetadata[any]) (uint16, string) {
	channel, filename, filedata, err := interpretArgs(meta.GetArgs())

	if err != nil {
		return 1, "[" + meta.GetSender().GetAddr() + "] Error processing SEND:" + err.Error()
	}

	fmt.Println("Received SEND command:", channel, filename, filedata)
	meta.GetSender().Send([]byte("received file"))
	return 0, ""
}
