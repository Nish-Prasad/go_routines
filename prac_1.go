package main

import "fmt"

func main() {
	unbuf := make(chan int)
	done := make(chan bool)

	go odd(unbuf, done)
	go even(unbuf, done)

	for i := 1; i <= 10; i++ {
		unbuf <- i
	}
	close(unbuf)

	<-done
	<-done
}

func odd(unbuf chan int, done chan bool) {
	for val := range unbuf {
		if val%2 != 0 {
			fmt.Println("Odd:", val)
		}
	}
	done <- true
}

func even(unbuf chan int, done chan bool) {
	for val := range unbuf {
		if val%2 == 0 {
			fmt.Println("Even:", val)
		}
	}
	done <- true
}