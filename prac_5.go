package main

import (
	"fmt"
)

/*
Write a Go program that uses two goroutines to print numbers from 1 to 10 alternately, where:

One goroutine prints odd numbers

The other prints even numbers

The output should be:
1 2 3 4 5 6 7 8 9 10
*/

func main() {

	oddChan := make(chan bool)
	evenChan := make(chan bool)
	done := make(chan bool)

	go printOdd(oddChan, evenChan, done)
	go printEven(oddChan, evenChan, done)

	oddChan <- true
	<-done
}

func printOdd(oddChan, evenChan, done chan bool) {

	for i := 1; i < 10; i = i + 2 {
		<-oddChan
		fmt.Println(i)
		evenChan <- true
	}

}

func printEven(oddChan, evenChan, done chan bool) {

	for i := 2; i <= 10; i = i + 2 {
		<-evenChan
		fmt.Println(i)
		if i == 10 {
			// close(oddChan)
			// close(evenChan)
			done <- true
			// close(done)
		} else {
			oddChan <- true
		}
	}
}
