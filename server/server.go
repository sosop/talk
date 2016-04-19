package server

import (
	"io"
	"net"
	"os"
)

type (
	Server struct {
		net   string
		laddr string
	}
)

func NewServer(net, laddr string) *Server {
	return &Server{net, laddr}
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
		go func(net.Conn) {
			defer conn.Close()
			io.Copy(os.Stdout, conn)
		}(conn)
	}
}
