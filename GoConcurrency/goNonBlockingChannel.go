package main

import "fmt"

func goNonBlockingChannel() {
	messages := make(chan string)
	signals := make(chan string)

	// receive of non-blocking
	// if messages have already values, then select is take with that values
	// But if messages don't have a values, then select is taken by default
	select {
	case msg := <-messages:
		fmt.Println("Message received", msg)
	default:
		fmt.Println("No message received")
	}

	// send of non-blocking
	// if messages include "buffer" and receiver, then select is taken by case messages<-msg
	// But if messages don't have a values, then select is taken by default
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no msg send")
	}

	select {
	case msg := <-messages:
		fmt.Println("Message received", msg)
	case sig := <-signals:
		fmt.Println("Signal received", sig)
	default:
		fmt.Println("no activity")
	}
}
