package server

import (
	"bufio"
	"chatroom-server/lists"
	"chatroom-server/logs"
	"container/list"
	"fmt"
	"log"
	"net"
)

var (
	clients  = list.New()
	message  = make(chan *Message)
	entrance = make(chan Client)
	exit     = make(chan Client)
)

func Start(port uint16) {
	addr := fmt.Sprint("0.0.0.0:", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	go handleChan()
	logs.Info("server started on port", port)
	defer listener.Close()
	for {
		accept(listener)
	}
}

func accept(listener net.Listener) {
	conn, err := listener.Accept()
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Debug(conn.LocalAddr(), conn.RemoteAddr())
	logs.Info("new connection from", conn.RemoteAddr())
	go receive(conn)
}

func receive(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	client := NewClient(scanner.Text(), conn)
	defer client.Close()
	entrance <- client
	message <- &Message{Content: client.GetName() + " enter", Sender: "server"}
	for scanner.Scan() {
		logs.Debug("recv:", scanner.Text())
		message <- &Message{Content: scanner.Text(), Sender: client.GetName()}
	}
	exit <- client
	message <- &Message{Content: client.GetName() + " exit", Sender: "server"}
}

func handleChan() {
	for {
		select {
		case msg := <-message:
			Publish(msg, clients)
			logs.Info(msg.Content, "published")
		case client := <-entrance:
			clients.PushBack(client)
			logs.Info(client.GetName(), "enter")
		case client := <-exit:
			lists.Remove(clients, client)
			logs.Info(client.GetName(), "exit")
		}
	}
}
