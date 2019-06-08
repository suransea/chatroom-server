package server

import (
	"chatroom-server/logs"
	"fmt"
	"net"
)

type Client struct {
	name string
	conn net.Conn
	ch   chan string
}

func NewClient(name string, conn net.Conn) *Client {
	client := new(Client)
	client.name = name
	client.conn = conn
	client.ch = make(chan string)
	go handleChan(client)
	return client
}

func (c *Client) Close() {
	c.conn.Close()
	close(c.ch)
}

func handleChan(client *Client) {
	for content := range client.ch {
		n, err := fmt.Fprintln(client.conn, content)
		if err != nil {
			logs.Error(err)
		}
		logs.Debug("send", n, "bytes to", client.conn.RemoteAddr())
		logs.Info("send", content, "to", client.name)
	}
}
