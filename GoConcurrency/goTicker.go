package main

import (
	"fmt"
	"time"
)

func goTicker() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at :", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("stop ticker")
}
