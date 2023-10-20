package main

import "fmt"

func IsPrimo(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func Primo(limit int, lista chan<- int) {
	start := 2
	for i := 0; i <= limit; i++ {
		for primo := start; ; primo++ {
			if IsPrimo(primo) {
				lista <- primo
				start = primo + 1
				break
			}
		}
	}
	close(lista)
}

func main() {
	lista := make(chan int)
	go Primo(1000, lista)
	for i := range lista {
		fmt.Printf(" %d", i)
	}
}
