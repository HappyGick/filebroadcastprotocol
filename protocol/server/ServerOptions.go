package server

type ServerOptions struct {
	MinPort    int
	MaxPort    int
	ListenPort int
	Hostname   string
}

func DefaultOptions() ServerOptions {
	return ServerOptions{
		32000,
		64000,
		5000,
		"localhost",
	}
}
