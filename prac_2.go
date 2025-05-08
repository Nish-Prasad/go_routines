package main

import (
	"fmt"
	"sync"
)

/*
Write a Go program that launches two goroutines:

One prints numbers from 1 to 5

The other prints letters from 'A' to 'E'
*/

func main() {
	numChan := make(chan bool)
	charChan := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2) // We have 2 goroutines

	go printNumbers(numChan, charChan, &wg)
	go printLetters(charChan, numChan, &wg)

	numChan <- true // Start by signaling the number printer

	wg.Wait() // Wait for both goroutines to finish
}

func printNumbers(numChan, charChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		<-numChan
		fmt.Println(i)
		charChan <- true
	}
}

func printLetters(charChan, numChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for ch := 'A'; ch <= 'E'; ch++ {
		<-charChan
		fmt.Println(string(ch))
		if ch != 'E' {
			numChan <- true
		}
	}
}
