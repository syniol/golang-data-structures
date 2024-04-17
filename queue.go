package data_structures

type Queue[T any] struct {
	elements []T
	size     int
}

type EmptyQueueError struct{}

func (EmptyQueueError) Error() string {
	return "queue is empty"
}

type Queuer[T any] interface {
	Enqueue(el T)
	Dequeue() (el T, err error)
	Peek() (el T, err error)
	Count() int
	Clear()
}

// NewQueue instantiates new instance of Queue Data Structure
func NewQueue[T any]() Queuer[T] {
	return &Queue[T]{
		elements: *new([]T),
		size:     0,
	}
}

// Enqueue Adds an object to the end of the Queue
func (q *Queue[T]) Enqueue(el T) {
	q.elements = append(q.elements, el)
	q.size = q.size + 1
}

// Dequeue Removes and returns the object at the beginning of the Queue
func (q *Queue[T]) Dequeue() (el T, err error) {
	if q.size == 0 {
		err = &EmptyQueueError{}

		return
	}

	el = q.elements[0]
	q.elements = q.elements[1:]
	q.size = q.size - 1

	return
}

// Peek Returns the object at the beginning of the Queue without removing it
func (q *Queue[T]) Peek() (el T, err error) {
	if q.size == 0 {
		err = &EmptyQueueError{}

		return
	}

	el = q.elements[0]

	return
}

// Count Gets the number of elements contained in the Queue
func (q *Queue[T]) Count() int {
	return q.size
}

// Clear Removes all objects from the Queue
func (q *Queue[T]) Clear() {
	q.elements = *new([]T)
	q.size = 0
}
