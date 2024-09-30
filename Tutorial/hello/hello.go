package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Gary")
	fmt.Println(message)

}
