package subtopics

import "fmt"

// Concurrency without the waitgroups

func SemaOne() {
	fmt.Println("\nSemaphore Example #1")
	chData := make(chan int)
	chDone := make(chan bool)

	// Create two go routines that generate data on chData
	go func() {
		for i := 0; i < 10; i++ {
			chData <- i
			fmt.Println("func 1: Tx ->", i)
		}
		chDone <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			chData <- i
			fmt.Println("func 2: Tx ->", i)
		}
		chDone <- true
	}()

	// This function won't close until it receives all the messages on chDone
	// It is a go routine because without a go routine running seperately, the
	// <- chDone statements will constantly block before the range statement
	// interrogates c!
	go func() {
		<-chDone
		<-chDone
		close(chData)
	}()

	// The range automatically blocks until chData is exhausted or closed
	for n := range chData {
		fmt.Printf("semaphore1: Rx <- %d\n", n)
	}

}
