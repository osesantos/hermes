package broker_test

import (
	"testing"

	"hermes/src/internal/broker"
)

func TestNewBroker(t *testing.T) {
	b := broker.NewBroker()

	if b == nil {
		t.Error("expected broker to be created")
	}
}

func TestBroker_Publish(t *testing.T) {
	b := broker.NewBroker()

	err := b.Publish("test", []byte("hello"))

	if !err.IsOk() {
		t.Errorf("expected no error; got %v", err)
	}
}
