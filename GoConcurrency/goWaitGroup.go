package main

import (
	"fmt"
	"sync"
	"time"
)

func worker1(id int) {
	fmt.Printf("Worker %v started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %v finished\n", id)

}

func goWaitGroup() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		//increments the counter by 1, indicating that one more goroutine (worker) will be added and that the program should wait for it to finish.
		wg.Add(1)

		i := i

		go func() {
			//defer schedules the wg.Done()
			defer wg.Done()
			worker1(i)
		}()
	}
	wg.Wait()
}
