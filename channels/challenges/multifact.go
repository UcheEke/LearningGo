package challenges

import (
	"math/rand"
	"fmt"
)

// MultiFact allows the solution of multiple factorial problems concurrently

func MultiFact(){
	// nGenerator: using n as the number of problems to solve, nGenerator
	// returns a channel which provides a random selection of n ints ranging
	// b/w 1 and 100
	nGenerator := func(n int) chan int {
		numCh := make(chan int)
		fmt.Printf("Generating %d psuedo-random numbers\n",n)
		go func() {for i := 0; i < n; i++ {
			numCh <- (1 + rand.Intn(13))
			}
			close(numCh)
		}()
		return numCh
	}

	// factorial: Calculates factorial of the input n
	factorial := func(n int) int {
		product := 1
		for i:=n; i>0; i-- {
			product *= i
		}
		return product
	}

	// doFactorials: generates all of the goroutines required, each producing a factorial
	doFactorials := func(in chan int) chan int {
		// Create the return channel
		out := make(chan int)
		// Critical design: Ensure that the created channel can be closed within the go routine
		go func() {
			for n := range in {
				out <- factorial(n)
			}
			close(out)
		}()
		// Return the created channel
		return out
	}

	// Blocking logic: ranging over factorials blocks main from exiting
	factorials := doFactorials(nGenerator(100))
	for solution := range factorials {
		fmt.Printf("%d\n",solution)
	}
}