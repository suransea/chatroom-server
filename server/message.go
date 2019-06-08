package server

type Message struct {
	Content string
	Sender  string
}

func NewMessage(content string, sender string) *Message {
	msg := new(Message)
	msg.Content = content
	msg.Sender = sender
	return msg
}

func Publish(msg *Message) {
	for e := clients.Front(); e != nil; e = e.Next() {
		str := msg.Sender + ": " + msg.Content
		e.Value.(*Client).ch <- str
	}
}
