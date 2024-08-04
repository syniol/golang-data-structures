package heap

import "golang.org/x/exp/constraints"

type Heap[T constraints.Ordered] struct {
	Nodes []T
}

func (h Heap[T]) Push(node T) {
	h.Nodes = append(h.Nodes, node)
}

func (h Heap[T]) Pop() T {
	return nil
}

func (h Heap[T]) Count() int {
	return 0
}
