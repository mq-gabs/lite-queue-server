package manager

import (
	"bytes"
	"lite_queue_server/protocol"
	"testing"
)

var m *QueueManager = New()
var id string

func TestNewQueue(t *testing.T) {
	if err := m.NewQueue("names"); err != nil {
		t.Fatalf("cannot create queue: %v", err)
	}

	if err := m.NewQueue("names"); err == nil {
		t.Fatal("it did not return error for create queue already existent")
	}
}

func TestPush(t *testing.T) {
	if err := m.Push("names", []byte("John Doe")); err != nil {
		t.Fatalf("cannot push: %v", err)
	}

	if _, err := m.getQueue("names"); err != nil {
		t.Fatalf("cannot get queue by name: %v", err)
	}
}

func TestPop(t *testing.T) {
	data, err := m.Pop("names")

	splitted := bytes.Split(data, []byte{protocol.Seperator})

	id = string(splitted[0])
	name := splitted[1]

	if err != nil {
		t.Fatalf("cannot pop: %v", err)
	}

	if string(name) != "John Doe" {
		t.Fatalf("returned value must be 'Johb Doe': %v", data)
	}
}

func TestAck(t *testing.T) {
	err := m.Ack("names", id)

	if err != nil {
		t.Fatalf("cannot ack: %v", err)
	}
}
