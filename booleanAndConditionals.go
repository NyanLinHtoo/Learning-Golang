package main

import "fmt"

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
