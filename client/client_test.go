package client

import (
	prot "talk/protocol"
	"testing"
)

var c = NewClient("tcp", "127.0.0.1:8008")

func TestConn(t *testing.T) {
	c.Conn()
	data := prot.Data{Text: []byte("hello")}
	p := prot.Protocol{From: "test1", Comm: prot.CONN, Data: data}
	c.Send(p)
	c.Quit.Wait()
	Close(c)
}
