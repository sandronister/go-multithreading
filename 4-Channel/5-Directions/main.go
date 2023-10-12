package main

import "fmt"

func receive(name string, hello chan<- string) {
	hello <- name
}

func read(data <-chan string) {
	fmt.Print(<-data)
}

func main() {
	hello := make(chan string)
	go receive("Hello", hello)
	read(hello)
}

/**
Receive only chan<-
Read only <-chan
*/
