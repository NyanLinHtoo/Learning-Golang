package main

import "fmt"

func channel() {
	// if make(chan string) => only one channel and need receiver to receive value
	messages := make(chan string)

	go func() { messages <- "Hello" }()

	// variable := <-channel => variable receive a value from a channel
	msg := <-messages

	fmt.Println(msg)
}
