package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// go mod edit -replace example.com/greetings=../greetings => point directory
	message := greetings.Hello("Gary")
	fmt.Println(message)

}
