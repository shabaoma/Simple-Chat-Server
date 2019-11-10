package message

import (
	"time"
)

type Message struct {
	From      string
	To        string
	Text      string
	CreatedAt time.Time
}

func New(from, to, text string) *Message {
	return &Message{from, to, text, time.Now()}
}
