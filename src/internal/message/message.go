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
