package main

import (
	"talk/server"
)

func main() {
	s := server.NewServer("tcp", "8008")
	s.ServeTCP()
}
