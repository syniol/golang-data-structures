package data_structures

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	sut := Queue[string]{
		elements: *new([]string),
		size:     0,
	}

	sut.Enqueue("syniol")
	sut.Enqueue("limited")

	if len(sut.elements) != 2 {
		t.Error("it was expecting two records in queue")
	}
}

func TestQueue_Peek(t *testing.T) {
	sut := Queue[string]{
		elements: *new([]string),
		size:     0,
	}

	_, err := sut.Peek()
	if err == nil {
		t.Error("expected an error")
	}
}
