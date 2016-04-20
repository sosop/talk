package main

import (
	"talk/server"
)

func main() {
	s := server.NewServerMemStore("tcp", ":8008")
	s.ServeTCP()
}
