package server

import (
	"fmt"
	"io"
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
		key := store(s, conn)
		go s.process(key, conn)
		fmt.Println(key + " connet success!")
	}
}

func store(s *Server, conn net.Conn) string {
	buf := make([]byte, 128)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		Logger.Error(err)
		return ""
	}
	p := protocol.UnSerializer(buf[:n])
	key := p.From
	s.Keep(key, conn)
	return key
}

func (s *Server) process(key string, conn net.Conn) {
	defer s.Del(key)
	defer conn.Close()
	count := 0
	flag := true
	for flag {
		buf := make([]byte, protocol.MaxDataSize)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				count++
				Logger.Error("empty data, conn may be disconnet!")
				if count == 10 {
					flag = false
				}
			} else {
				Logger.Error(err)
			}
			continue
		}
		if n == 0 {
			Logger.Error("data too large")
		}
		p := protocol.UnSerializer(buf[:n])
	}
}
