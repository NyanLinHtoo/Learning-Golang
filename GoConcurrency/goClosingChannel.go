package main

import "fmt"

func goClosingChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// j,more := <- jobs is special 2 value receive. And if jobs channel is closed and all values are received, more value is turned into false.
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				//if all jobs are successfully received, done will be notify
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job ", j)
	}
	close(jobs) // for channel close, if close is not include, deadlock occur because of jobs channel is 5 and j is only have 3
	fmt.Println("sent all job")

	<-done // this line is block before worker notify on channel
}
