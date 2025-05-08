package main

import (
	"fmt"
	"sync"
)

/*
Write a Go program that launches three goroutines:

One prints "Ping"

One prints "Pong"

One prints the iteration number (1 to 5)
*/

func main() {

	numChan := make(chan bool)
	pingChan := make(chan bool)
	pongChan := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(3)
	go printNum(numChan, pingChan, pongChan, &wg)
	go printPing(numChan, pingChan, pongChan, &wg)
	go printPong(numChan, pingChan, pongChan, &wg)

	numChan <- true

	wg.Wait()

}

func printNum(numChan, pingChan, pongChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		<-numChan
		fmt.Println(i)
		pingChan <- true
	}

}

func printPing(numChan, pingChan, pongChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		<-pingChan
		fmt.Println("ping")
		pongChan <- true
	}
}

func printPong(numChan, pingChan, pongChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		<-pongChan
		fmt.Println("pong")
		if i != 5 {
			numChan <- true
		}
	}
}
