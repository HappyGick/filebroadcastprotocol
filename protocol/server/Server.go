package server

import (
	"fmt"
	"net"
	"os"

	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
	"github.com/HappyGick/filebroadcastprotocol/protocol/server/commands"
)

type HandleConnectionFunction func(conn common.Connection)

type FileBroadcastServer struct {
	Channels    map[int]Channel
	CommHandler ServerCommandHandler
	MinPort     int
	MaxPort     int
	host        string
	listenPort  int
}

func Create(options ServerOptions) FileBroadcastServer {
	s := FileBroadcastServer{
		map[int]Channel{},
		NewServerCommandHandler(),
		options.MinPort,
		options.MaxPort,
		options.Hostname,
		options.ListenPort,
	}

	s.setupCommands()

	return s
}

func (s FileBroadcastServer) handleConnectionDefault(conn *common.Connection) {
	defer conn.Close()
	defer fmt.Println("Successfully disconnected from", conn.GetAddr())
	fmt.Println("Successfully connected to", conn.GetAddr())
	for {
		data, err := conn.Receive()
		if err != nil {
			fmt.Println("Error receiving data from", conn.GetAddr(), ":", err)
			return
		}
		err = s.CommHandler.Handle(data, conn)
		if err != nil {
			fmt.Println("Error processing command from", conn.GetAddr(), ":", err)
			return
		}
	}
}

func (s FileBroadcastServer) setupCommands() {
	s.CommHandler.RegisterCommand(commands.NewSendToChannelCommand())
}

func (s FileBroadcastServer) Listen() {
	l, err := net.Listen("tcp", s.host+":"+fmt.Sprint(s.listenPort))

	if err != nil {
		fmt.Println("Error on listen:", err)
		os.Exit(1)
	}

	defer l.Close()

	fmt.Println("Listening on", s.host+":"+fmt.Sprint(s.listenPort))

	for {
		c, err := l.Accept()

		if err != nil {
			fmt.Println("Error on connection:", err)
			continue
		}

		conn := common.NewConnection(&c)

		go s.handleConnectionDefault(&conn)
	}
}
