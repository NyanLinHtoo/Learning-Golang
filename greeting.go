package main

import "fmt"

var points = []int{30, 52, 14, 56, 77}

func sayGreet(n string) {
	fmt.Printf("Hello %v", n)
}

func showScore() {
	fmt.Printf("You scored %0.2f", score)
}
