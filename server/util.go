package server

import "container/list"

func Publish(msg *Message, clients *list.List) {
	for e := clients.Front(); e != nil; e = e.Next() {
		str := msg.Sender + ": " + msg.Content
		e.Value.(*Client).Send(str)
	}
}
