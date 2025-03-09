package message_test

import (
	"testing"

	"hermes/src/internal/message"

	"github.com/google/uuid"
)

func TestNewMessage(t *testing.T) {
	message := message.NewMessage("test", []byte("hello"))

	if message.ID == uuid.Nil {
		t.Error("expected message ID to be set")
	}

	if message.Topic != "test" {
		t.Errorf("expected message topic to be 'test'; got %s", message.Topic)
	}

	if string(message.Payload) != "hello" {
		t.Errorf("expected message payload to be 'hello'; got %s", message.Payload)
	}

	if message.Timestamp.IsZero() {
		t.Error("expected message timestamp to be set")
	}
}

func TestMessage_String(t *testing.T) {
	message := message.NewMessage("test", []byte("hello"))

	if message.String() != "hello" {
		t.Errorf("expected message string to be 'hello'; got %s", message.String())
	}
}
