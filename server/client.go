package server

import (
	"chatroom-server/logs"
	"fmt"
	"net"
)

type Client interface {
	GetName() string
	Send(content string)
	Close()
}

type client struct {
	name string
	conn net.Conn
	send chan string
}

func NewClient(name string, conn net.Conn) Client {
	c := new(client)
	c.name = name
	c.conn = conn
	c.send = make(chan string)
	go handleSend(c)
	return c
}

func (c *client) GetName() string {
	return c.name
}

func (c *client) Send(content string) {
	c.send <- content
}

func (c *client) Close() {
	c.conn.Close()
	close(c.send)
}

func handleSend(c *client) {
	for content := range c.send {
		n, err := fmt.Fprintln(c.conn, content)
		if err != nil {
			logs.Error(err)
		}
		logs.Debug("send", n, "bytes to", c.conn.RemoteAddr())
		logs.Info("send", content, "to", c.name)
	}
}
