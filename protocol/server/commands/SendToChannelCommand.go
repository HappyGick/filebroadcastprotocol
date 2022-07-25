package commands

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
	"github.com/HappyGick/filebroadcastprotocol/protocol/common/util"
)

func interpretArgs(args []byte) (uint64, string, []byte, error) {
	r := bytes.NewReader(args)
	channel := make([]byte, 8)
	var filename bytes.Buffer
	w := bufio.NewWriter(&filename)
	size := make([]byte, 8)

	_, err := r.Read(channel)

	if err != nil {
		return 0, "", nil, err
	}

	buf, err := util.ReadUntil(r, 0)

	if err != nil {
		return 0, "", nil, err
	}

	w.Write(buf)
	w.Flush()

	_, err = r.Read(size)

	if err != nil {
		return 0, "", nil, err
	}

	data := make([]byte, binary.BigEndian.Uint64(size))

	_, err = r.Read(data)

	if err != nil {
		return 0, "", nil, err
	}

	return binary.BigEndian.Uint64(channel), filename.String(), data, nil
}

func executeCommand(meta common.CommandMetadata) (uint16, string) {
	channel, filename, filedata, err := interpretArgs(meta.GetArgs())

	if err != nil {
		return 1, err.Error()
	}

	fmt.Println("Received SEND command:", channel, filename, filedata)
	meta.GetSender().Send([]byte("received file"))
	return 0, ""
}

func NewSendToChannelCommand() ServerCommand {
	return ServerCommand{
		Name:            "SEND",
		ExecuteFunction: executeCommand,
	}
}
