package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

type fnumber struct {
	number int
	factorial int
}

func factorial(ch <-chan int) <-chan fnumber {
	out := make(chan fnumber)
	go func(){
		var fn fnumber
		number := <-ch
		product := 1
		for i := number; i > 0; i-- {
			product *= i
		}
		fn.number = number
		fn.factorial = product
		out <- fn
		close(out)
	}()
	return out
}

func numberGenerator(n int) <-chan int {
	// Seed the random number generator with current time in nanoseconds
	rand.Seed(int64(time.Now().Nanosecond()))
	// Create a channel to deliver the random numbers on
	out := make(chan int)
	// local func randomNumbers
	randomNumbers := func(limit int) {
		for i := 0; i < limit; i++ {
			out <- (1 + rand.Intn(20))
		}
		close(out)
	}
	go randomNumbers(n)
	return out
}

func fanOutFactorials(in <-chan int, limit int) []<-chan fnumber {
	out := make([]<-chan fnumber, 0, limit)
	for i := 0; i < limit; i++ {
		out = append(out, factorial(in))
	}
	return out
}

func fanInResults(chans ...<-chan fnumber) <-chan fnumber {
	var wg sync.WaitGroup
	out := make(chan fnumber)

	wg.Add(len(chans))

	results := func(ch <-chan fnumber){
		for n := range ch {
			out <- n
		}
		wg.Done()
	}

	for _, ch := range chans {
		go results(ch)
	}

	go func(){
		wg.Wait()
		close(out)
	}()

	return out
}

func main(){
	limit := 10000
	fnSlice := fanInResults(fanOutFactorials(numberGenerator(limit),limit)...)
	var counter int
	var n fnumber
	for n = range fnSlice {
		counter++
		fmt.Printf("%d: %d! = %d\n",counter, n.number,n.factorial)
	}
}

