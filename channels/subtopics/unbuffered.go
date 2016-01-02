package subtopics

import (
	"fmt"
	"sync"
	"time"
)

func Unbuffered(n int, wg sync.WaitGroup) {
	fmt.Println("\nUnbuffered channel demonstration...")
	c := make(chan int)
	wg.Add(2)
	go func(){
		for i:= 0; i < n; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func(){
		for {
			fmt.Printf("Unbuffered: Data Rx <- %d\n", <- c)
		}
		wg.Done()
	}()

	time.Sleep(time.Duration(1 * time.Second))
	fmt.Println("Waiting for a second or two....")

}
