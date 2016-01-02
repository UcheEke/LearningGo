package subtopics

// Illustrates the passing of channels as function arguments
import "fmt"

// PassingChannels creates a generator channel on one goroutine
// and a processor channel in another.
func PassingChannels() {
	fmt.Println("\nPassing Channels example...")

	// Tx integers from 0-9 and then close
	generator := func() <-chan int {
		outCh := make(chan int)

		go func() {
			for i := 0; i < 10; i++ {
				fmt.Printf("generator: Tx -> %d\n", i)
				outCh <- i
			}
			close(outCh)
		}()
		return outCh
	}

	// Rx integers from generator channel and sum. Send resulting channel to containing function
	receiver := func(inCh <-chan int) chan int {
		outCh := make(chan int)
		// This go routine sums the output from inCh
		go func() {
			var sum int
			for data := range inCh {
				fmt.Printf("receiver: Rx <- %d\n", data)
				sum += data
			}
			outCh <- sum
			close(outCh)
		}()
		return outCh
	}

	genCh := generator()
	sumCh := receiver(genCh)

	// Rx from sumCh and exit
	for n := range sumCh {
		fmt.Printf("PassingChannels: Rx <- %d\n", n)
	}
}
