package main

import (
	"fmt"
	"sort"
	"strings"
)

func standardLibrary() {
	// strings packages (doesn't affact on original)
	greeting := "Hello there friends"

	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, " "))

	// original string (greeting)
	fmt.Println("The original string is :", greeting)

	// Sort Package (change original value)
	ages := []int{20, 65, 77, 12, 85, 32, 2, 14, 90}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 32) // if number is exist,return index of this number.If number isn't exist, return index to insert number
	fmt.Println(index)

	names := []string{"nyan", "lin", "htoo", "furtive"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "nyan"))
}
