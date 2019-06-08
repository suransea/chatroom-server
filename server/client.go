package server

import (
	"chatroom-server/logs"
	"fmt"
	"net"
)

type Client struct {
	Name string
	Conn net.Conn
	Chan chan string
}

func NewClient(name string, conn net.Conn) *Client {
	client := new(Client)
	client.Name = name
	client.Conn = conn
	client.Chan = make(chan string)
	go client.handleChan()
	return client
}

func (c *Client) Send(content string) {
	c.Chan <- content
}

func (c *Client) Close() {
	c.Conn.Close()
	close(c.Chan)
}

func (c *Client) handleChan() {
	for content := range c.Chan {
		n, err := fmt.Fprintln(c.Conn, content)
		if err != nil {
			logs.Error(err)
		}
		logs.Debug("send", n, "bytes to", c.Conn.RemoteAddr())
		logs.Info("send", content, "to", c.Name)
	}
}
