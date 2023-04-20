package test

import (
	"github.com/AKushch1337/architecture-lab3/painter"
	"golang.org/x/exp/shiny/screen"
	"testing"
)

type mockOperation struct{}

func (m mockOperation) Do(t screen.Texture) (ready bool) {
	return false
}

func TestMessageQueue_push(t *testing.T) {
	mq := &painter.MessageQueue{}

	op := &mockOperation{}
	mq.Push(op)
	if len(mq.Queue) != 1 {
		t.Errorf("Expected queue length to be 1, but got %d", len(mq.Queue))
	}
	if mq.Queue[0] != op {
		t.Error("Expected pushed operation to be in the queue")
	}

	op2 := &mockOperation{}
	mq.Push(op2)
	if len(mq.Queue) != 2 {
		t.Errorf("Expected queue length to be 2, but got %d", len(mq.Queue))
	}
	if mq.Queue[1] != op2 {
		t.Error("Expected pushed operation to be in the queue")
	}
}

func TestMessageQueue_pull(t *testing.T) {
	mq := &painter.MessageQueue{}

	op := mq.Pull()
	if op != nil {
		t.Error("Expected nil operation when pulling from an empty queue")
	}

	op2 := &mockOperation{}
	mq.Push(op2)
	pulledOp := mq.Pull()
	if len(mq.Queue) != 0 {
		t.Errorf("Expected queue length to be 0, but got %d", len(mq.Queue))
	}
	if pulledOp != op2 {
		t.Error("Expected pulled operation to be the same as the pushed operation")
	}

	op3 := &mockOperation{}
	op4 := &mockOperation{}
	mq.Push(op3)
	mq.Push(op4)
	pulledOp2 := mq.Pull()
	if len(mq.Queue) != 1 {
		t.Errorf("Expected queue length to be 1, but got %d", len(mq.Queue))
	}
	if pulledOp2 != op3 {
		t.Error("Expected pulled operation to be the first pushed operation")
	}
}
