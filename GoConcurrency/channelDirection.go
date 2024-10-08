package main

import "fmt"

// send value to channel
func ping(pings chan string, msg string) {
	pings <- msg
}

// receive value
func pong(pings chan string, pongs chan string) {
	msg := <-pings
	pongs <- msg
}

func channelDirection() {
	pings := make(chan string, 1) // make channel with 1 time receive
	pongs := make(chan string, 1) // make channel with 1 time receive
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
