# Queue Data Structure in Golang
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam


## Encapsulation
In this example you could see `NewQueue`, produces a `struct` matching `Queuer[T any]` interface.
This is by design to avoid accessing `elements` and `size` of the queue. However, `NewQueue` is not 
utilised inside tests to examine the actual value for `size` and elements in the queue.

```go
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

func NewQueue[T any]() Queuer[T] {
    return &Queue[T]{
        items: *new([]T),
        size:  0,
    }
}
```


### Tests

```bash
go test ./queue/... -v
```


##### Credits
Author: [Hadi Tajallaei](mailto:hadi@syniol.com)

Copyright &copy; 2024, Syniol Limited. All rights reserved.
