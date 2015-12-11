//Tests for Queue

package queue

import (
	"testing"
)

func TestEmptyQueuePopsNil(t *testing.T) {
	subject := Queue{}
	item, err := subject.Pop()
	if item != nil {
		t.Errorf("Expected item to equal nil, got: %v", item)
	}
	if err == nil {
		t.Errorf("Expected err 'No Items In Queue', got: %v", err)
	}
}

func TestNonEmptyQueuePopsItem(t *testing.T) {
	subject := Queue{}
	subject.Push("test")
	item, err := subject.Pop()
	if item == nil {
		t.Errorf("Expected item to  NOT equal nil, got: %v", item)
	}
	if item != "test" {
		t.Errorf("Expected item to equal 'test', got: %v", item)
	}
	if err != nil {
		t.Errorf("Expect err to equal nil, go: %v", err)
	}
}

func TestPopItemRemovesItFromTheFrontOfTheQueue(t *testing.T) {
	subject := Queue{}
	subject.Push("test")
	subject.Pop()
	itemNil, _ := subject.Pop()
	if itemNil != nil {
		t.Errorf("Expected Item to equal nil got %v", itemNil)
	}
}

func TestMultiplePushAndPop(t *testing.T) {
	subject := Queue{}
	subject.Push("test")
	subject.Push("second")
	subject.Pop()
	secondItem, _ := subject.Pop()
	if secondItem != "second" {
		t.Errorf("Expected Item to equal 'second' got %v", secondItem)
	}
}

func TestCanGetLengthOfQueue(t *testing.T) {
	subject := Queue{}
	length := subject.Length()
	if length != 0 {
		t.Errorf("Expected length to equal 0, got %d", length)
	}
	subject.Push("test")
	subject.Push("second")
	length = subject.Length()
	if length != 2 {
		t.Errorf("Expected length to equal 2, got %d", length)
	}
	subject.Pop()
	length = subject.Length()
	if length != 1 {
		t.Errorf("Expected length to equal 1, got %d", length)
	}
}
