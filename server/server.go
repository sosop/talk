package server

import "net"

type (
	Server struct {
		net   string
		laddr string
		conns map[string]net.Conn
	}
)

func NewServer(network, laddr string) *Server {
	conns := make(map[string]net.Conn, 1024)
	return &Server{network, laddr, conns}
}

func (s *Server) ServeTCP() {
	l, err := net.Listen(s.net, s.laddr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			// TODO
		}

	}
}
