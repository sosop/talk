package server

import "testing"

var s *Server

func TestServerTCP(t *testing.T) {
	s = NewServer("tcp", ":8008")
	s.ServeTCP()
}
