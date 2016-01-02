package subtopics

import "fmt"

func One2Many() {

	fmt.Println("\nOne2Many example...")
	numProcs := 10
	dataCh := make(chan int)
	doneCh := make(chan bool)

	generator := func(size int) {
		for i := 0; i < size; i++ {
			fmt.Printf("generator: Tx -> %d\n", i)
			dataCh <- i
		}
		close(dataCh)
	}

	go generator(100)

	for i := 0; i < numProcs; i++ {
		go func(a int) {
			for n := range dataCh { // range ends when dataCh closes.

				fmt.Printf("goroutine #%d: Rx <- %d\n", a, n)

			}
			doneCh <- true // don't forget to tell the 'blocking logic' when you're done!!
		}(i)
	}

	// Blocking logic (waiting on doneCh semaphores)
	for i := 0; i < numProcs; i++ {
		<-doneCh
	}
}
