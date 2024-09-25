package data_structures

import (
	"fmt"
)

//var sss map[string]string
//
//
//
//ssss := new(map[string]string)
//
//aaaa := make(map[string]string, 10)

func ExampleMapCreationWithRawAssignment() {
	var mapUnderTest map[string]string

	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Println(err)
		}
	}()

	mapUnderTest["company"] = "Syniol Limited"

	// Output:
	// assignment to entry in nil map
}

func ExampleMapCreationWithAssignment() {
	var mapUnderTest map[string]string

	mapUnderTest = map[string]string{
		"location": "London",
	}

	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Println(err)
		}
	}()

	mapUnderTest["company"] = "Syniol Limited"

	fmt.Println(mapUnderTest)

	// Output:
	// map[company:Syniol Limited location:London]
}

func ExampleMapCreationWith() {

}
