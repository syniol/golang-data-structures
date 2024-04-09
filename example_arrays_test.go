package data_structures_test

import (
	"fmt"
)

// Ordered
// Static size
// Arrays are values by default (meaning passing a copy of variable in a function)

func ExampleArrayOfNumbers() {
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
