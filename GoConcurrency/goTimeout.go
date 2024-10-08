package main

import (
	"fmt"
	"time"
)

func goTimeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Result one"
	}()

	// In select, case res and case <- time .... will work simultaneously and
	// in case res => c1 will take 2 seconds to send value to res and
	// in case <- time ... will take 1 seconds
	// select will take 1st ready to receive, So "timeout one" will be printed out
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("TimeOut one")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result two"
	}()

	// In select, case res and case <- time .... will work simultaneously and
	// in case res => c2 will take 2 seconds to send value to res and
	// in case <- time ... will take 3 seconds
	// select will take 1st ready to receive, So "result two" will be printed out
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout two")
	}
}
