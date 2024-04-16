package data_structures

type Queue[T any] struct {
	elements []T
	size     int
}

type EmptyQueueError struct{}

func (EmptyQueueError) Error() string {
	return "queue is empty"
}

func (q *Queue[T]) Enqueue(el T) {
	q.elements = append(q.elements, el)
	q.size = q.size + 1
}

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

func (q *Queue[T]) Peek() (el T, err error) {
	if q.size == 0 {
		err = &EmptyQueueError{}

		return
	}

	el = q.elements[0]

	return
}

func (q *Queue[T]) Count() int {
	return q.size
}

func (q *Queue[T]) Clear() {
	q.elements = *new([]T)
	q.size = 0
}
