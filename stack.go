package data_structures

type Stack[T any] struct {
	elements []T
	count    int
}

type StackEmptyError struct{}

func (StackEmptyError) Error() string {
	return "stack is empty"
}

type Stacker[T any] interface {
	Push(el T)
	Peek() (el T, err error)
	Pop() (el T, err error)
	Count() int
}

// NewStack instantiates new instance of Stack Data Structure
func NewStack[T any]() Stacker[T] {
	return &Stack[T]{
		elements: *new([]T),
		count:    0,
	}
}

// Push Inserts an object at the top of the Stack.
func (s *Stack[T]) Push(el T) {
	if s.count == 0 {
		s.elements = append(s.elements, el)
		s.count = 1

		return
	}

	currentHead := s.elements[0]

	s.elements[0] = el
	s.elements[s.count] = currentHead

	s.count = s.count + 1
}

// Peek Returns the object at the top of the Stack without removing it.
func (s *Stack[T]) Peek() (el T, err error) {
	if s.count == 0 {
		err = &StackEmptyError{}

		return
	}

	el = s.elements[0]

	return
}

// Pop Removes and returns the object at the top of the Stack
func (s *Stack[T]) Pop() (el T, err error) {
	if s.count == 0 {
		err = &StackEmptyError{}

		return
	}

	el = s.elements[0]
	s.elements = s.elements[1:]

	s.count = s.count - 1

	return
}

// Count returns current number of elements inside the stack
func (s *Stack[T]) Count() int {
	return s.count
}
