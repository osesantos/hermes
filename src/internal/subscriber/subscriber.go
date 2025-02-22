package subscriber

import (
	"hermes/src/internal/message"

	"github.com/google/uuid"
)

type Subscriber struct {
	ID      uuid.UUID
	Channel chan message.Message
}
