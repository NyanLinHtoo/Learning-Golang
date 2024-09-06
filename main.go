package main

import "fmt"

func main() {
	// variables()
	// printFormatString()
	arrays()
}

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

func printFormatString() {
	age := 26
	name := "Nyan"

	// Print
	fmt.Print("My age is ", age, "\n")
	fmt.Print("My name is ", name, "\n")

	// Println
	fmt.Println("My age is ", age)
	fmt.Println("My name is ", name)

	// Printf => (formatted Strings)
	fmt.Printf("My age is %v and my name is %v \n", age, name) // %v => value
	fmt.Printf("My age is %q and my name is %q \n", age, name) // %q => quotes arround variable(only string)
	fmt.Printf("My age type is %T  \n", age)                   // %T => Type
	fmt.Printf("My score is %f  \n", 255.647)                  // %f => float
	fmt.Printf("My score is %0.1f  \n", 255.647)               // %0.1f => 1 decimal place(rounded)
	fmt.Printf("My score is %0.2f  \n", 255.647)               // %0.2f => 2 decimal place(rounded)

	// Sprintf (save formatted Strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name) // stored in a variable
	fmt.Println("The save string: ", str)
}

func arrays() {}
