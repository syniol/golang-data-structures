package data_structures

import (
	"encoding/json"
	"fmt"
	"sort"
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

func ExampleMapCreationWithMakeAndLength() {
	mapUnderTest := make(map[string]string, 1)

	mapUnderTest["company"] = "Syniol Limited"
	mapUnderTest["location"] = "London"

	fmt.Println(mapUnderTest)

	// Output:
	// map[company:Syniol Limited location:London]
}

func ExampleMapAccessingExistingElement() {
	mapUnderTest := map[string]any{
		"name":     "Syniol Limited",
		"location": "London",
	}

	fmt.Println(mapUnderTest["location"])

	// Output:
	// London
}

func ExampleMapAccessingNonExistentElement() {
	mapUnderTest := map[string]any{
		"name":     "Syniol Limited",
		"location": "London",
	}
	// catch or finally
	defer func() {
		// catch
		if err := recover(); err == nil {
			fmt.Println("no error detected for nil pointer")
		}
	}()

	fmt.Println(mapUnderTest["website"])

	// Output:
	// <nil>
	// no error detected for nil pointer
}

func ExampleMapToJSON() {
	mapUnderTest := map[string]any{
		"name": "Syniol Limited",
	}

	jsonMap, _ := json.Marshal(mapUnderTest)

	fmt.Println(string(jsonMap))

	// Output:
	// {"name":"Syniol Limited"}
}

func ExampleMapWithoutAnyElementIteration() {
	mapUnderTest := make(map[string]any)

	for k, v := range mapUnderTest {
		fmt.Println(k)
		fmt.Println(v)
	}

	// Output:
	//
}

func ExampleMapWithElementsIteration() {
	mapUnderTest := make(map[string]any)
	mapUnderTest["name"] = "Syniol Limited"
	mapUnderTest["location"] = "London"

	for k, v := range mapUnderTest {
		fmt.Println(k, ":", v)
	}

	// Output:
	// name : Syniol Limited
	// location : London
}

func ExampleMapWithElementsTraditionalIteration() {
	mapUnderTest := make(map[int]any)
	mapUnderTest[0] = "Syniol Limited"
	mapUnderTest[1] = "London"

	for i := 0; i < len(mapUnderTest); i++ {
		fmt.Println(i, ":", mapUnderTest[i])
	}

	// Output:
	// 0 : Syniol Limited
	// 1 : London
}

func ExampleMapSorting() {
	mapUnderTest := map[string]any{
		"name":     "Syniol Limited",
		"location": "London",
	}

	keys := make([]string, 0, len(mapUnderTest))
	for k := range mapUnderTest {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, ":", mapUnderTest[k])
	}

	// Output:
	// location : London
	// name : Syniol Limited
}

func ExampleMapElementChangeWithoutPointerWithValue() {
	mapUnderTest := map[string]any{
		"name":     "Syniol Limited",
		"location": "London",
	}

	func(val map[string]any) {
		val["location"] = "Miami"
		val["website"] = "https://syniol.com"
	}(mapUnderTest)

	fmt.Println(mapUnderTest["location"])
	fmt.Println("Length of map is:", len(mapUnderTest))

	// Output:
	// Miami
	// Length of map is: 3
}

func ExampleMapDeletingExistingElement() {
	mapUnderTest := map[string]any{
		"name":     "Syniol Limited",
		"location": "London",
		"website":  "https://syniol.com",
	}

	delete(mapUnderTest, "location")

	fmt.Println(mapUnderTest)

	// Output:
	// map[name:Syniol Limited website:https://syniol.com]
}

func ExampleMapDeletingNonExistentElement() {
	mapUnderTest := map[string]any{
		"name":     "Syniol Limited",
		"location": "London",
		"website":  "https://syniol.com",
	}

	// catch or finally
	defer func() {
		// catch
		if err := recover(); err == nil {
			fmt.Println("no error detected for nil pointer")
		}
	}()

	delete(mapUnderTest, "url")

	// Output:
	// no error detected for nil pointer
}
