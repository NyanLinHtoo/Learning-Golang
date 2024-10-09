package main

import (
	"fmt"
	"time"
)

// if we want to run in the specified time use timer and ticker
func goTimers() {
	time1 := time.NewTimer(2 * time.Second)

	<-time1.C // C mean channel => C <- time.Time
	fmt.Println("Timer 1 fired")

	// time2 has enough time to run but "stop2 := time2.Stop()" is stop time2.
	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
