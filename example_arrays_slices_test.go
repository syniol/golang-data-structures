package data_structures_test

import (
	"fmt"
)

func ExampleArraySizeAndDefinition() {
	var numbers [10]int

	cities := [3]string{
		"London",
		"Manchester",
		"Birmingham",
	}

	fmt.Println(numbers)
	fmt.Println(cities)

	// Output:
	// [0 0 0 0 0 0 0 0 0 0]
	// [London Manchester Birmingham]
}

func ExampleArrayInitialisationWithNew() {
	var numbers *[10]int = new([10]int)

	fmt.Println(numbers)
	fmt.Println(*numbers)
	fmt.Println("length:", len(*numbers))
	fmt.Println("capacity:", cap(*numbers))

	// Output:
	// &[0 0 0 0 0 0 0 0 0 0]
	// [0 0 0 0 0 0 0 0 0 0]
	// length: 10
	// capacity: 10
}

func ExampleArrayInitialisationWithMakeAndLength() {
	var numbers []int = make([]int, 10)
	fmt.Println(numbers)
	fmt.Println("length:", len(numbers))
	fmt.Println("capacity:", cap(numbers))

	// Output:
	// [0 0 0 0 0 0 0 0 0 0]
	// length: 10
	// capacity: 10
}

func ExampleArrayInitialisationWithMakeAndLengthAndCapacity() {
	var numbers []int = make([]int, 10, 25)

	fmt.Println(numbers)
	fmt.Println("length:", len(numbers))
	fmt.Println("capacity:", cap(numbers))

	// Output:
	// [0 0 0 0 0 0 0 0 0 0]
	// length: 10
	// capacity: 25
}

func ExampleSliceSizeAndDefinition() {
	var numbers []int

	cities := []string{
		"London",
		"Manchester",
		"Birmingham",
	}

	fmt.Println(numbers)
	fmt.Println(cities)

	// Output:
	// []
	// [London Manchester Birmingham]
}

func ExampleSliceInitialisationWithNew() {
	var numbers *[]int = new([]int)

	fmt.Println(numbers)
	fmt.Println(*numbers)
	fmt.Println("length:", len(*numbers))
	fmt.Println("capacity:", cap(*numbers))

	// Output:
	// &[]
	// []
	// length: 0
	// capacity: 0
}

func ExampleModifyingArrayAndSlice() {
	numbers := []int{1, 2, 3}
	numbers[2] = 2000

	fmt.Println(numbers)
	// Output:
	// [1 2 2000]
}

func ExampleModifyingArrayAndSliceWithAppend() {
	numbers := []int{1, 2, 3}

	numbers = append(numbers, 4)
	numbers = append(numbers, []int{5, 6}...)
	numbers = append(numbers, []int{7, 8}[1:]...)

	fmt.Println(numbers)
	// Output:
	// [1 2 3 4 5 6 8]
}
