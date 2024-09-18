package main

import "fmt"

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
