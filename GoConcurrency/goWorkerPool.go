package main

import (
	"fmt"
	"time"
)

func workers(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Workers", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Workers", id, "finished job", j)
		results <- j * 2
	}
}

func goWorkerPool() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go workers(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j

	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
