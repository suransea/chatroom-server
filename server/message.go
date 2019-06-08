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
