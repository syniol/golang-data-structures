package queue

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	sut := queue[string]{
		items: *new([]string),
		size:  0,
	}

	sut.Enqueue("syniol")
	sut.Enqueue("limited")

	if len(sut.items) != 2 {
		t.Error("it was expecting two records in queue")
	}
}

func TestQueue_Dequeue(t *testing.T) {
	sut := queue[string]{
		items: *new([]string),
		size:  0,
	}

	const firstItemInQueue = "syniol"

	sut.Enqueue(firstItemInQueue)
	sut.Enqueue("limited")

	el, err := sut.Dequeue()
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

	if len(sut.items) != 1 {
		t.Error("it was expecting a single record in queue")
	}
}

func TestQueue_Peek(t *testing.T) {
	sut := queue[string]{
		items: *new([]string),
		size:  0,
	}

	_, err := sut.Peek()
	if err == nil {
		t.Error("expected an error")
	}

	firstItemInQueue := "syniol"
	sut.Enqueue(firstItemInQueue)
	sut.Enqueue("limited")

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

	if len(sut.items) != 2 {
		t.Error("it was expecting all records in the queue after peek")
	}
}

func TestQueue_Clear(t *testing.T) {
	sut := queue[string]{
		items: *new([]string),
		size:  0,
	}

	sut.Enqueue("syniol")
	sut.Enqueue("limited")

	if len(sut.items) != 2 {
		t.Error("it was expecting two records in queue")
	}

	sut.Clear()

	if len(sut.items) != 0 {
		t.Error("it was expecting no record in the queue")
	}
}

func TestQueue_Count(t *testing.T) {
	sut := queue[string]{
		items: *new([]string),
		size:  0,
	}

	if sut.Count() != 0 {
		t.Error("it was expecting no record in the queue")
	}

	sut.Enqueue("syniol")
	sut.Enqueue("limited")

	if sut.Count() != 2 {
		t.Error("it was expecting two records in the queue")
	}
}
