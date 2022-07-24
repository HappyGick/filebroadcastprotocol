package server

import (
	"fmt"
	"net"
	"os"

	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
)

type HandleConnectionFunction func(conn common.Connection)

type FileBroadcastServer struct {
	Channels    map[int]Channel
	CommHandler CommandHandler
	MinPort     int
	MaxPort     int
	host        string
	listenPort  int
}

func Create(options ServerOptions) FileBroadcastServer {
	return FileBroadcastServer{
		map[int]Channel{},
		CommandHandler{},
		options.MinPort,
		options.MaxPort,
		options.Hostname,
		options.ListenPort,
	}
}

func (s FileBroadcastServer) handleConnectionDefault(conn common.Connection, handle HandleConnectionFunction) {
	defer fmt.Println("Successfully disconnected from ", conn.GetAddr())
	fmt.Println("Successfully connected to ", conn.GetAddr())
	handle(conn)
}

func (s FileBroadcastServer) Listen(handle HandleConnectionFunction) {
	s.ListenWithPort(handle, s.listenPort)
}

func (s FileBroadcastServer) ListenWithPort(handle HandleConnectionFunction, port int) {
	l, err := net.Listen("tcp", s.host+":"+fmt.Sprint(port))

	if err != nil {
		fmt.Println("Error on listen: ", err)
		os.Exit(1)
	}

	defer l.Close()

	fmt.Println("Listening on ", s.host+":"+fmt.Sprint(s.listenPort))

	for {
		c, err := l.Accept()

		if err != nil {
			fmt.Println("Error on connection: ", err)
			continue
		}

		conn := common.NewConnection(&c)

		go s.handleConnectionDefault(conn, handle)
	}
}
