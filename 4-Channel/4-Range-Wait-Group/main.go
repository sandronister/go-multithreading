package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	ch := make(chan int)
	go publish(ch)
	go reader(ch, &wg)
	wg.Wait()
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d \n", x)
		wg.Done()
	}
}
