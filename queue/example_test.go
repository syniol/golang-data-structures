package queue_test

import (
	"data-structures/queue"
	"fmt"
)

type QueueElement struct {
	Body string
}

func ExampleNewQueue() {
	systemUnderTest := queue.NewQueue[*QueueElement]()
	systemUnderTest.Enqueue(&QueueElement{
		Body: "{}",
	})

	fmt.Println(systemUnderTest.Count())

	// Output:
	// 1
}
