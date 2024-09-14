package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

var score = 99.556

func main() {
	// variables()
	// printFormatString()
	// arraysAndSlices()
	// standardLibrary()
	// loops()
	// booleanAndConditionals()

	// Using Function
	// sayGreeting("nyan")
	// sayGreeting("lin")
	// sayBye("htoo")
	// cycleNames([]string{"Nyan", "Lin", "Htoo"}, sayGreeting)
	// cycleNames([]string{"Nyan", "Lin", "Htoo"}, sayBye)
	// a1 := circleArea(10.34)
	// a2 := circleArea(58.235)
	// fmt.Println(a1, a2)
	// fmt.Printf("Circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)

	// Return Multiple Values
	// fn1, sn1 := getInitials("nick thomas")
	// fmt.Println(fn1, sn1)

	// fn2, sn2 := getInitials("Juliet thomas")
	// fmt.Println(fn2, sn2)

	// fn3, sn3 := getInitials("Monyoe")
	// fmt.Println(fn3, sn3)

	// Packeage Scope
	// sayGreet("nyan")
	// for _, v := range points {
	// 	fmt.Println(v)
	// }
	// showScore()

	maps()
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

func loops() {
	x := 0
	for x < 5 {
		fmt.Println("the value of x : ", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("the value of i : ", i)
	}

	names := []string{"nyan", "lin", "htoo", "furtive"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("The value at index is %v and %v position \n", index, value)
	}

	// if we don't need index, use "underscore"
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
		value := "new String" // nothing happened because "value" is local copy of variable
		fmt.Println(value)
	}
}

func booleanAndConditionals() {
	age := 45
	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")

	} else {
		fmt.Println("Age is no less than 40")
	}

	names := [4]string{"nyan", "lin", "htoo", "furtive"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
}

// one argument
func sayGreeting(n string) {
	fmt.Printf("Good Morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("GoodBye %v \n", n)
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// Two argument
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

// Multiple Return Values
func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

// Map
func maps() {
	// String map
	menu := map[string]float64{
		"samsung": 12.66,
		"iphone":  15.66,
		"Oppo":    10.32,
		"vivo":    11.11,
	}

	fmt.Println(menu)
	fmt.Println(menu["vivo"])

	// Looping Map
	for key, value := range menu {
		fmt.Println(key, "- ", value)
	}

	// Integer Map
	phoneNumber := map[int]string{
		792826217: "Nyan",
		792826218: "Lin",
		792826219: "Htoo",
	}

	fmt.Println(phoneNumber)
	fmt.Println(phoneNumber[792826217])

	// Changing Value in Map
	phoneNumber[792826218] = "LinLin"
	fmt.Println(phoneNumber)

	phoneNumber[792826219] = "HtooHtoo"
	fmt.Println(phoneNumber)
}
