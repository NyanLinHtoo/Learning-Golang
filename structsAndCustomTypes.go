package main

import (
	"fmt"
)

// declare Custom Types
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// Receiver Function with pointer
// Format the bill
func (b *bill) format() string {
	fs := "Bill Breakdown : \n"
	var total float64 = 0

	//items list
	for k, v := range b.items {
		fs += fmt.Sprintf("%-12v ...$%v \n", k+":", v)
		total += v
	}

	//Tip
	fs += fmt.Sprintf("%-12v ...$%0.2f \n", "Tip:", b.tip)

	//Total
	fs += fmt.Sprintf("%-12v ...$%0.2f \n", "Total:", total)

	return fs
}

// update Tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// Add Items
func (b *bill) addItems(name string, price float64) {
	b.items[name] = price
}
