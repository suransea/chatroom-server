package main

import (
	"chatroom-server/conf"
	"chatroom-server/server"
)

func main() {
	server.Start(conf.Port)
}
