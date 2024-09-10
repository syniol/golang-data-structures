package stack

import "testing"

func TestQueue_Push(t *testing.T) {
	sut := stack[string]{
		elements: *new([]string),
		size:     0,
	}

	sut.Push("syniol")
	sut.Push("limited")

	if len(sut.elements) != 2 {
		t.Error("it was expecting two records in queue")
	}
}

func TestQueue_Pop(t *testing.T) {
	sut := stack[string]{
		elements: *new([]string),
		size:     0,
	}

	const firstItemInQueue = "syniol"

	sut.Push(firstItemInQueue)
	sut.Push("limited")

	el, err := sut.Pop()
	if err != nil {
		t.Errorf(
			"it was not expecting an error but got: %s",
			err.Error(),
		)
	}

	if el != firstItemInQueue {
		t.Errorf(
			"it was expecting the first element '%s' in a queue to be dequeued",
			firstItemInQueue,
		)
	}

	if len(sut.elements) != 1 {
		t.Error("it was expecting a single record in queue")
	}
}

func TestQueue_Pop_Error(t *testing.T) {
	sut := stack[string]{
		elements: *new([]string),
		size:     0,
	}

	_, err := sut.Pop()
	if err == nil {
		t.Error("it was expecting an error")
	}

	if err.Error() != "queue is empty" {
		t.Error("it was expecting an error reporting an empty queue")
	}
}

func TestQueue_Peek(t *testing.T) {
	sut := stack[string]{
		elements: *new([]string),
		size:     0,
	}

	_, err := sut.Peek()
	if err == nil {
		t.Error("expected an error")
	}

	firstItemInQueue := "syniol"
	sut.Push(firstItemInQueue)
	sut.Push("limited")

	item, errPeak := sut.Peek()
	if errPeak != nil {
		t.Errorf(
			"it was not expecting an error but got: %s",
			errPeak.Error(),
		)
	}

	if item != firstItemInQueue {
		t.Errorf(
			"it was expecting the first element '%s' in a queue to be returned",
			firstItemInQueue,
		)
	}

	if len(sut.elements) != 2 {
		t.Error("it was expecting all records in the queue after peek")
	}
}

func TestQueue_Count(t *testing.T) {
	sut := stack[string]{
		elements: *new([]string),
		size:     0,
	}

	if sut.Count() != 0 {
		t.Error("it was expecting no record in the queue")
	}

	sut.Push("syniol")
	sut.Push("limited")

	if sut.Count() != 2 {
		t.Error("it was expecting two records in the queue")
	}
}
