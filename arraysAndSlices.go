package main

import "fmt"

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
