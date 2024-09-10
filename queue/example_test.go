package queue_test

import (
	"fmt"

	"data-structures/queue"
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
