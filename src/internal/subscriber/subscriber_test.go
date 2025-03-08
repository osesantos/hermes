package subscriber

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewSubscriber(t *testing.T) {
	subscriber := NewSubscriber()

	if subscriber.ID == (uuid.UUID{}) {
		t.Error("expected subscriber ID to be set")
	}

	if cap(subscriber.Channel) != 10 {
		t.Errorf("expected subscriber channel capacity to be 10; got %d", cap(subscriber.Channel))
	}

	if subscriber.Topic != "" {
		t.Errorf("expected subscriber topic to be empty; got %s", subscriber.Topic)
	}
}

func TestSubscriber_Subscribe(t *testing.T) {
	subscriber := NewSubscriber()

	subscriber.Subscribe("test")

	if cap(subscriber.Channel) != 10 {
		t.Errorf("expected subscriber channel capacity to be 10; got %d", cap(subscriber.Channel))
	}

	if subscriber.Topic != "test" {
		t.Errorf("expected subscriber topic to be 'test'; got %s", subscriber.Topic)
	}
}
