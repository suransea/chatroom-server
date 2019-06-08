package server

type Message struct {
	Content string
	Sender  string
}

func NewMessage(content string, sender string) *Message {
	m := new(Message)
	m.Content = content
	m.Sender = sender
	return m
}
