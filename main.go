package main

import "fmt"

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

	// maps()

	// Pass by Values
	// passByValues()

	// Pointer
	// pointer()

	// Structs and Custom Types
	// myBill := newBill("Nyan's Bill")
	// fmt.Println(myBill)

	// Receiver Function
	myBill := newBill("Nyan's Bill")
	fmt.Println(myBill.format())
}
