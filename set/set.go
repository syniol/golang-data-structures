package set

type Set struct {
	elements map[any]bool
	size     int
}

func NewSet() *Set {
	return &Set{
		elements: make(map[any]bool),
		size:     0,
	}
}

func (s *Set) Add(el any) {
	s.elements[el] = true
	s.size = s.size + 1
}

func (s *Set) Remove(el any) {
	delete(s.elements, el)
	s.size = s.size - 1
}

func (s *Set) Size() int {
	return s.size
}

func (s *Set) IsEmpty() bool {
	return s.size == 0
}

func (s *Set) List() (list []any) {
	for el, _ := range s.elements {
		list = append(list, el)
	}

	return
}

func (s *Set) Has(el any) bool {
	_, exists := s.elements[el]
	if exists {
		return true
	}

	return false
}
