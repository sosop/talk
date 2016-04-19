package server

import (
	"net"
	"talk/protocol"
)

type (
	Server struct {
		net   string
		laddr string
		Store
	}
)

func NewServerMemStore(network, laddr string) *Server {
	return NewServer(network, laddr, NewMemoryStore())
}

func NewServer(network, laddr string, store Store) *Server {
	return &Server{network, laddr, store}
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
		go store(s, conn)
	}
}

func store(s *Server, conn net.Conn) {
	buf := make([]byte, 0, protocol.MaxDataSize)
	_, err := conn.Read(buf)
	if err != nil {
		// TODO
	}
	p := protocol.UnSerializer(buf)
	s.Store.Keep(p.From, conn)
}
