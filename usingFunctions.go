package main

import (
	"fmt"
	"math"
)

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
