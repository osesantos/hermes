package subscriber

import (
	"hermes/src/internal/message"

	"github.com/google/uuid"
)

type Subscriber struct {
	ID      uuid.UUID
	Channel chan message.Message
	Topic   string
}

func NewSubscriber() Subscriber {
	return Subscriber{
		ID:      uuid.New(),
		Channel: make(chan message.Message, 10),
		Topic:   "",
	}
}

func (s *Subscriber) Subscribe(topic string) {
	s.Channel = make(chan message.Message, 10)
	s.Topic = topic
}

func (s *Subscriber) Unsubscribe() {
	close(s.Channel)
	s.Topic = ""
}
