package subtopics

import (
	"fmt"
)

// Very much like the PassingChannels example but using two channels as generators and two
// accumulators to sum the generator outputs. The final statement requires the output from
// two unbuffered channels which will block the function from exiting before they operate

func DualAccumulator(limit int) {

	fmt.Println("\nDual Accumulator Example...")

	generator := func(id int) <-chan int {
		out := make(chan int)
		go func() {
			for i := 0; i < limit; i++ {
				fmt.Printf("gen%d: Tx -> %d\n", id,1)
				out <- 1
			}
			close(out)
		}()
		return out
	}

	receiver := func(in <-chan int, id int) chan int {
		out := make(chan int)
		go func() {
			var sum int
			for data := range in {
				fmt.Printf("rec%d: Rx <- %d\n", id, data)
				sum += data
			}
			fmt.Printf("rec%d: Tx -> %d (sum)\nClosing channel...\n", id, sum)
			out <- sum
		}()
		return out
	}

	gen1 := generator(1)
	gen2 := generator(2)
	acc1 := receiver(gen1,1)
	acc2 := receiver(gen2,2)
	fmt.Printf("DualAccumulator: Rx <- %d\n\n", <-acc1+<-acc2)
}
