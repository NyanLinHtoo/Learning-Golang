package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// variables()
	// printFormatString()
	// arraysAndSlices()
	standardLibrary()
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

func arraysAndSlices() {
	// Arrays
	// var ages [3]int = [3]int{100, 20, 52}
	var ages = [3]int{100, 20, 52}
	names := [4]string{"nyan", "lin", "htoo", "furtive"}
	names[2] = "HtooHtoo"
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices (use arrays under the hood)
	var scores = []int{100, 20, 25}
	scores[1] = 86
	scores = append(scores, 222)
	fmt.Println(scores, len(scores))

	//Slice Range
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	rangeOne = append(rangeOne, "NyanLinHtoo")
	fmt.Println(rangeOne)
}

func standardLibrary() {
	// strings packages (doesn't affact on original)
	greeting := "Hello there friends"

	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, " "))

	// original string (greeting)
	fmt.Println("The original string is :", greeting)

	// Sort Package (change original value)
	ages := []int{20, 65, 77, 12, 85, 32, 2, 14, 90}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 32) // if number is exist,retunr index of this number.If number isn't exist, return index to insert number
	fmt.Println(index)

	names := []string{"nyan", "lin", "htoo", "furtive"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "nyan"))
}
