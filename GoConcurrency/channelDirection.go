package main

import "fmt"

// send value to channel (chan <-)
func ping(pings chan string, msg string) {
	pings <- msg
}

// receive value from channel ( <- chan)
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
