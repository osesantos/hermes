package broker

import (
	"sync"
	"time"

	"hermes/src/internal"
	"hermes/src/internal/message"
	"hermes/src/internal/subscriber"

	"github.com/google/uuid"
)

type Broker struct {
	mu          sync.RWMutex
	subscribers map[string][]subscriber.Subscriber
	topics      map[string][]message.Message
}

// NewBroker creates a new Broker instance.
func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]subscriber.Subscriber),
		topics:      make(map[string][]message.Message),
	}
}

// Publish publishes a message to a topic.
func (b *Broker) Publish(topic string, payload []byte) internal.Result[error] {
	b.mu.Lock()
	defer b.mu.Unlock()

	msg := message.Message{
		ID:        uuid.New(),
		Topic:     topic,
		Payload:   payload,
		Timestamp: time.Now(),
	}

	b.topics[topic] = append(b.topics[topic], msg)

	if subscribers, exists := b.subscribers[topic]; exists {
		for _, subscriber := range subscribers {
			select {
			case subscriber.Channel <- msg: // Send the message to the subscriber.
			default: // If channel is full, skip the subscriber.
			}
		}
	}

	return internal.Success[error](nil)
}
