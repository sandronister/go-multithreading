package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() {
		channel <- "Hello World"
	}()

	msg := <-channel

	fmt.Println(msg)
}
