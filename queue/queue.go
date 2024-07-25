package queue

type Queue[T any] struct {
	items []T
	size  int
}

type Queuer[T any] interface {
	Enqueue(item T)
	Dequeue() (item T, err error)
	Peek() (item T, err error)
	Count() int
	Clear()
}

type ErrorEmptyQueue struct{}

func (ErrorEmptyQueue) Error() string {
	return "queue is empty"
}

// NewQueue Instantiates new instance of Queue Data Structure
func NewQueue[T any]() Queuer[T] {
	return &Queue[T]{
		items: *new([]T),
		size:  0,
	}
}

// Enqueue Adds an object to the end of the Queue
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
	q.size = q.size + 1
}

// Dequeue Removes and returns the object at the beginning of the Queue
func (q *Queue[T]) Dequeue() (item T, err error) {
	if q.size == 0 {
		err = &ErrorEmptyQueue{}

		return
	}

	item = q.items[0]
	q.items = q.items[1:]
	q.size = q.size - 1

	return
}

// Peek Returns the object at the beginning of the Queue without removing it
func (q *Queue[T]) Peek() (item T, err error) {
	if q.size == 0 {
		err = &ErrorEmptyQueue{}

		return
	}

	item = q.items[0]

	return
}

// Count Gets the number of items contained in the Queue
func (q *Queue[T]) Count() int {
	return q.size
}

// Clear Removes all objects from the Queue
func (q *Queue[T]) Clear() {
	q.items = *new([]T)
	q.size = 0
}
