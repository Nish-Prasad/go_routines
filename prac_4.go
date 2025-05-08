package main

/*
Write a Go program that starts two goroutines:

One goroutine generates even numbers from 2 to 10.

The other goroutine generates odd numbers from 1 to 9.
*/

import (
	"fmt"
	"sync"
)

func main() {
	numChan := make(chan int)

	var wg sync.WaitGroup         // for all goroutines (2 producers + 1 consumer)
	var producerWg sync.WaitGroup // only for 2 producers

	wg.Add(3)
	producerWg.Add(2)

	go produceEvens(numChan, &wg, &producerWg)
	go produceOdds(numChan, &wg, &producerWg)
	go consume(numChan, &wg, &producerWg)

	wg.Wait()
}

func produceEvens(out chan int, wg *sync.WaitGroup, producerWg *sync.WaitGroup) {
	defer wg.Done()
	defer producerWg.Done()

	for i := 2; i <= 10; i += 2 {
		out <- i
	}
}

func produceOdds(out chan int, wg *sync.WaitGroup, producerWg *sync.WaitGroup) {
	defer wg.Done()
	defer producerWg.Done()

	for i := 1; i <= 9; i += 2 {
		out <- i
	}
}

func consume(in chan int, wg *sync.WaitGroup, producerWg *sync.WaitGroup) {
	defer wg.Done()

	// Only waits for the producers
	go func() {
		producerWg.Wait()
		close(in)
	}()

	for val := range in {
		fmt.Println(val)
	}
}
