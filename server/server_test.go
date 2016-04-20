package server

import "testing"

var s *Server

func TestServerTCP(t *testing.T) {
	s = NewServerMemStore("tcp", ":8008")
	s.ServeTCP()
}
