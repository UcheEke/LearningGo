package subtopics

import (
	"fmt"
	"sync"
)

func N2oneRace(wg sync.WaitGroup) {
	// Sets up 3 go routines talking on the same channel

	fmt.Println("\nN-to-1 example...")
	c := make(chan int)
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("func 1: Tx ->", i)
			c <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("func 2: Tx ->", i)
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		fmt.Println("Closing channel...")
		close(c)
	}()

	for n := range c { // This automatically blocks!
		fmt.Printf("n2one: Rx <- %d\n", n)
	}
}
