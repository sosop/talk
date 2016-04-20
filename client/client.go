package client

import (
	"io/ioutil"
	"net"
	"sync"
	"talk/protocol"
	. "talk/slog"
	"time"
)

type (
	Client struct {
		net     string
		addr    string
		timeout time.Duration
		conn    net.Conn
		newMu   sync.Mutex
		closeMu sync.Mutex
		wg      sync.WaitGroup
	}
)

func NewClient(net, addr string) *Client {
	return &Client{net: net, addr: addr, timeout: time.Second * 6}
}

func (c *Client) Conn() {
	defer c.newMu.Unlock()
	c.newMu.Lock()
	var err error
	if c.conn == nil {
		c.conn, err = net.DialTimeout(c.net, c.addr, c.timeout)
		if err != nil {
			panic(err)
		}
	}
}

func (c *Client) Send(pro protocol.Protocol) {
	defer c.wg.Done()
	c.wg.Add(1)
	data := protocol.Serializer(pro)
	_, err := c.conn.Write(data)
	if err != nil {
		Logger.Error(err)
	}
}

func (c *Client) Read() protocol.Protocol {
	buf, err := ioutil.ReadAll(c.conn)
	if err != nil {
		Logger.Error(err)
	}
	return protocol.UnSerializer(buf)
}

func Close(c *Client) {
	c.wg.Wait()
	defer c.closeMu.Unlock()
	c.closeMu.Lock()
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
}
