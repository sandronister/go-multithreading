package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(20 * time.Millisecond)
	}
}

func main() {
	data := make(chan int)
	qtdWorks := 1000

	for i := 0; i <= qtdWorks; i++ {
		go worker(i, data)
	}

	for i := 0; i <= 1000000000; i++ {
		data <- i
	}
}
