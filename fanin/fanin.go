package main

/*

	Illustration of a simple generator/worker pattern with concurrency
	Generator sends out a stream of random numbers.

	A series of 10 workers pick up the next available number on the channel and process it
	(calculating the log2 in this case)

 */

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func generator(outCh chan float64) {
	for {
		rand.Seed(int64(time.Now().Nanosecond()))
		number := float64(rand.Int63n(int64(1000)))
		outCh <- number
	}
}

func worker(inCh chan float64, outCh chan float64, id int) {
	for {
		number := <-inCh
		result := math.Log2(number)
		fmt.Printf("Worker %d: Log2(%.2f) = %.2f\n", id, number, result)
		outCh <- result
	}
}

func main() {
	numbersCh := make(chan float64)
	processCh := make(chan float64)

	defer close(numbersCh)
	defer close(processCh)

	go generator(numbersCh)
	for i := 0; i < 10; i++ {
		fmt.Printf("Creating worker #%d\n", i)
		go worker(numbersCh, processCh, i)
	}

	var count int
	for {
		select {
		case <-processCh:
			count++
			fmt.Printf("numbers processed: %d\r", count)
			time.Sleep(300 * time.Millisecond)
		}
	}
}
