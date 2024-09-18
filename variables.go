package main

import "fmt"

func variables() {
	// Strings
	var nameOne string = "Nyan"
	var nameTwo = "Lin"
	nameThree := "Htoo" //shorthand
	fmt.Println(nameOne, nameTwo, nameThree)

	// Integer
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40 // shorthand
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and momery
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint16 = 21546 //unsigned Integer (contain zero or positive integer)
	fmt.Println(numOne, numTwo, numThree)

	// Float
	var scoreOne float32 = 25.9
	scoreTwo := 685865.3325 // shorthand (default float64)
	fmt.Println(scoreOne, scoreTwo)
}
