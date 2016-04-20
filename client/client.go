package client

import (
	"io/ioutil"
	"net"
	"sync"
	"talk/protocol"
	. "talk/slog"
	"time"
)

var (
	count = 0
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
		Quit    sync.WaitGroup
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
		c.Quit.Add(1)
		go c.process()
	}
}

func (c *Client) process() {
	defer c.Quit.Done()
	for c.conn != nil {
		c.Read()
	}
}

func (c *Client) Send(pro protocol.Protocol) error {
	defer c.wg.Done()
	c.wg.Add(1)
	data := protocol.Serializer(pro)
	_, err := c.conn.Write(data)
	if err != nil {
		Logger.Error(err)
	}
	return err
}

func (c *Client) Read() protocol.Protocol {
	buf, err := ioutil.ReadAll(c.conn)
	if err != nil {
		Logger.Error(err)
		return protocol.Protocol{}
	}
	if len(buf) == 0 {
		count++
		Logger.Error("empty data, conn may be disconnet!")
		if count == 10 {
			Close(c)
			c.Conn()
		}
		return protocol.Protocol{}
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
