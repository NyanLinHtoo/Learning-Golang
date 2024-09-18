package main

import "fmt"

// declare Custom Types
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"cake":   5.66,
			"coffee": 9.36,
			"coke":   6.13,
		},
		tip: 0,
	}

	return b
}

// Format the bill
// Receiver Function
func (b bill) format() string {
	fs := "Bill Breakdown : \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-12v ...$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-12v ...$%0.2f", "Total:", total)

	return fs
}
