# Stack Data Structure in Golang



## Generics
You will require to use a Go version that has the generic implementation. This allowed me to ensure, this code could
be utilised in various applications where type of elements inside the stack could be programmed.

```go
type StackElement struct {
    Body string
}

myStack := stack.NewStack[*StackElement]()
myStack.Push(&StackElement{
    Body: "{}",
})
```

you could find an example of this inside `example_test.go`.


## Processing & Memory Optimisation
You might have noticed the use of size instead of the use `len`; the built-in library in Go,
this is due to CPU and Memory utilisation during the calculation when stack exceeds a high number.


## Encapsulation
In this example you could see `NewStack`, produces a private `struct` matching `Queuer[T any]` interface.
This is by design to avoid accessing `items` and `size` of the stack. However, `NewStack` is not
utilised inside tests to examine the actual value for `size` and elements in the stack.


__Stack Struct__
```go
type stack[T any] struct {
    elements []T
    size  int
}
```

__Stacker Interface__
```go
type Stacker[T any] interface {
    Push(el T)
    Peek() (el T, err error)
    Pop() (el T, err error)
    Count() int
}
```

__NewStack__
```go
func NewStack[T any]() Queuer[T] {
    return &stack[T]{
        elements: *new([]T),
        size:  0,
    }
}
```


### Tests
From the root of the repository, please run the following command:

```bash
go test ./stack/... -v
```


##### Credits
Author: [Hadi Tajallaei](mailto:hadi@syniol.com)

Copyright &copy; 2024, Syniol Limited. All rights reserved.
