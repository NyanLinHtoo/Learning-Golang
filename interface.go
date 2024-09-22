package main

import (
	"fmt"
	"math"
)

// Interfaces
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

// Square Method
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// Circle Method
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is : %0.2f \n", s, s.area())
	fmt.Printf("CircumFerence of %T is : %0.2f \n", s, s.circumf())
}
