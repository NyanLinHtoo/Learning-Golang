package main

import "fmt"

func updateName(x string) string {
	x = "Special Name"
	return x //to change value
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 150.22
}

func passByValues() {
	// group A types (Non Pointer values) => strings, ints, floats, booleans, arrays, structs
	name := "Normal Name"

	// nothing changing original name("Normal Name") because updateName(name) changed the copy of name variable
	// updateName(name)
	// fmt.Println(name)

	// changed values because func updateName
	name = updateName(name)
	fmt.Println(name)

	// Group B Types (Pointer Wrapper Values) => slices, maps,functions
	menu := map[string]float64{
		"pie":       36.44,
		"ice-cream": 25.66,
	}

	// Changed values
	updateMenu(menu)
	fmt.Println(menu)

}
