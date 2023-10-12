package main

import "fmt"

func main() {
	ch := make(chan int)
	go publish(ch)
	reader(ch)
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d \n", x)
	}
}
