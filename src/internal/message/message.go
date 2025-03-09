package message

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID
	Topic     string
	Payload   []byte
	Timestamp time.Time
}

func NewMessage(topic string, payload []byte) Message {
	return Message{
		ID:        uuid.New(),
		Topic:     topic,
		Payload:   payload,
		Timestamp: time.Now(),
	}
}

func (m Message) String() string {
	return string(m.Payload)
}
