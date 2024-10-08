package main

import "fmt"

func goRangeOverChannel() {
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"

	close(queue) // channel is close and then iterate will stop after received 3 values

	for elem := range queue {
		fmt.Println(elem)
	}
}
