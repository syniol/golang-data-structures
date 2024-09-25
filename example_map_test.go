package data_structures

import (
	"fmt"
)

func ExampleMapCreationWithRawAssignment() {
	var mapUnderTest map[string]string

	// catch or finally
	defer func() {
		// catch
		if err := recover(); err != nil {
			fmt.Println("error", err)
		}
	}()

	mapUnderTest["company"] = "Syniol Limited"

	// Output:
	// error assignment to entry in nil map
}

func ExampleMapCreationWithAssignment() {
	mapUnderTest := map[string]string{
		"location": "London",
	}

	// catch or finally
	defer func() {
		// catch
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	mapUnderTest["company"] = "Syniol Limited"

	fmt.Println(mapUnderTest)

	// Output:
	// map[company:Syniol Limited location:London]
}

func ExampleMapCreationWithNew() {
	mapUnderTest := *new(map[string]string)

	// catch or finally
	defer func() {
		// catch
		if err := recover(); err != nil {
			fmt.Println("error", err)
		}
	}()

	mapUnderTest["company"] = "Syniol Limited"

	// Output:
	// error assignment to entry in nil map
}

func ExampleMapCreationWithMake() {
	mapUnderTest := make(map[string]string)

	// catch or finally
	defer func() {
		// catch
		if err := recover(); err != nil {
			fmt.Println("error", err)
		}
	}()

	mapUnderTest["company"] = "Syniol Limited"

	fmt.Println(mapUnderTest)

	// Output:
	// map[company:Syniol Limited]
}
