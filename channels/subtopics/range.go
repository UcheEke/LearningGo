package subtopics

import "fmt"

func BlockingWithRange(n int) {
	fmt.Println("\nBlocking with range keyword demonstration...")
	c := make(chan int)

	go func() {
		for i := 1; i <= n; i++ {
			c <- i
		}
		close(c)
	}()

	for r := range c {
		fmt.Printf("blockWithRange: channel -> %d\n", r)
	}
}
