package subtopics

import (
	"fmt"
)

func SemaTwo() {
	// Similar to SemaOne(), only this time using for loops to generate go routines
	// and another for loop to generate an equal number of <-done 'waits'

	chData := make(chan int)
	chDone := make(chan bool)
	var numProc int = 10

	fmt.Println("\nSemaphore Example #2...")
	// Generate multiple go routines
	for i := 0; i < numProc; i++ {
		go func(a int) {
			for j := 0; j < 3; j++ {
				fmt.Printf("go routine #%d Tx -> %d\n", a, j)
				chData <- j
			}
			chDone <- true
		}(i)
	}

	// Generate an equal number of wait processes
	go func() {
		for i := 0; i < numProc; i++ {
			<-chDone
			fmt.Printf("Closed chDone #%d\n", i)
		}
		close(chData)
		close(chDone)
	}()

	for n := range chData {
		fmt.Printf("Semaphore2: Rx <- %d\n", n)
	}
}
