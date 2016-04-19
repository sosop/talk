package client

import (
	"net"
	"sync"
	"time"
)

type (
	Client struct {
		net     string
		addr    string
		timeout time.Duration
		conn    net.Conn
		mu      sync.Mutex
		wg      sync.WaitGroup
	}
)

func NewClient(net, addr string) *Client {
	return &Client{net: net, addr: addr, timeout: time.Second * 6}
}

func (c *Client) Conn() {
	c.mu.Lock()
	defer c.mu.Unlock()
	var err error
	c.conn, err = net.DialTimeout(c.net, c.addr, c.timeout)
	if err != nil {
		panic(err)
	}
}

func (c *Client) Send(msg string) {
	defer c.wg.Done()
	c.wg.Add(1)
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		// TODO
	}
}

func (c *Client) Read() string {
	defer c.wg.Done()
	c.wg.Add(1)
	buf := make([]byte, 0, 2048)
	_, err := c.conn.Read(buf)
	if err != nil {
		// TODO
	}
	return string(buf)
}

func Close(c *Client) {
	defer c.conn.Close()
	c.wg.Wait()
}
