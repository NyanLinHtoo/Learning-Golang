package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Print("done")

	done <- true //send value for notifying work done
}

func channelSynchronization() {
	done := make(chan bool, 1)
	go worker(done)

	// if "<-done" is not include, program will done before start worker
	<-done // wait
}
