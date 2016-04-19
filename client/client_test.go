package client

import "testing"

var c = NewClient("tcp", "127.0.0.1:8008")

func TestConn(t *testing.T) {
	c.Conn()
	c.Send("hello")
	Close(c)
}
