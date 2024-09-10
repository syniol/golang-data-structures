package stack_test

import (
	"fmt"

	"data-structures/stack"
)

type StackElement struct {
	Body string
}

func ExampleNewStack() {
	systemUnderTest := stack.NewStack[*StackElement]()
	systemUnderTest.Push(&StackElement{
		Body: "{}",
	})

	fmt.Println(systemUnderTest.Count())

	// Output:
	// 1
}
