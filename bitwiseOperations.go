package main

import "fmt"

func bitwiseOperations() {
	a := 12 // 1100 in binary
	b := 10 // 1010 in binary

	// AND ( & ) // set each bit to 1 if both bits are 1
	fmt.Println("AND Operation =>", a&b) // 1000 // 8

	// OR ( | ) // set eact bit to 1 if either of bits is 1
	fmt.Println("OR Operation =>", a|b) // 1110 // 14

	// XOR ( ^ ) // set each bit to 1 if only one of two bit is 1
	fmt.Println("XOR Operation =>", a^b) // 0110 // 6

	// AND NOT ( &^ ) // clear(set to 0) all bits that are set in right operand

	// 	How AND NOT (&^) works:
	// For each bit position:

	// If the bit in the right operand is 1, the corresponding bit in the left operand is set to 0.
	// If the bit in the right operand is 0, the corresponding bit in the left operand remains unchanged.
	fmt.Println("AND NOT Operation =>", a&^b) // 0100 // 4

	// Left Shift ( << 2 ) => 2 is for moving positions
	fmt.Println("Left Shift Operation =>", a<<3) // 0000 1100 => 1100000 // 96
	fmt.Println("Left Shift Operation =>", b<<3) // 0000 1010 => 1010000 // 80

	// Right Shift ( >> 2 ) => 2 is for moving positions
	fmt.Println("Right Shift Operation =>", a>>3) // 0000 1100 => 0000 0001 // 1
	fmt.Println("Right Shift Operation =>", b>>3) // 0000 1010 => 0000 0001 // 1

}
