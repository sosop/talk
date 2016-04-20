package server

import (
	"fmt"
	"io/ioutil"
	"net"
	"talk/protocol"
	. "talk/slog"
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
			Logger.Error(err)
		}
		store(s, conn)
	}
}

func store(s *Server, conn net.Conn) {
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		Logger.Error(err)
		return
	}
	p := protocol.UnSerializer(buf)
	fmt.Println(p)
	s.Store.Keep(p.From, conn)
}
