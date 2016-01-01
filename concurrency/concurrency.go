package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Create a global waitGroup
var wg sync.WaitGroup

// Creating a random sleep timer to introduce some variablity into each function call
// leading to some interleaving in the outputs
func randomSleep() {
	time.Sleep(time.Duration(rand.Intn(30) * int(time.Millisecond)))
}

// After Go 1.5 this init() function is no longer necessary since parallelism is the default setting.
// This ensured that parallelism occurred in Go pre 1.5
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// Functions foo and bar can be run concurrently using global waitGroup wg
func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo:", i)
		randomSleep()
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar:", i)
		randomSleep()
	}
	wg.Done()
}

// Check for race conditions by using the -race option when using the cmdline 'go run' command
func race_condition() {
	var counter int
	var wgRc sync.WaitGroup

	fmt.Println("Creating race condition...")

	increment := func(s string) {
		var x int
		for i := 0; i < 5; i++ {
			x = counter
			fmt.Printf("%s: Counter = %d\n", s, counter)
			randomSleep()
			x++
			counter = x
		}
		wgRc.Done()
	}
	wgRc.Add(2)
	go increment("foo")
	go increment("bar")
	wgRc.Wait()
}

// The race condition function can be modified with a mutex in order to lock the
// counter variable in between calls. Mutex stops other processes from modifying a
// variable between its Lock and Unlock functions
func useMutex() {
	var counter int
	var wgRc sync.WaitGroup
	var mtx sync.Mutex

	fmt.Println("Eliminating race condition via mutex...")

	increment := func(s string) {
		for i := 0; i < 5; i++ {
			randomSleep()
			{
				mtx.Lock()
				counter++
				fmt.Printf("%s: Counter = %d\n", s, counter)
				mtx.Unlock()
			}
		 }
		wgRc.Done()
	}
	wgRc.Add(2)
	go increment("foo")
	go increment("bar")
	wgRc.Wait()
}

func atomicVariables(){
	// Defining a variable as atomic acts in a similar manner to using a mutex
	var aCount int64
	var wgA sync.WaitGroup
	fmt.Printf("\nTesting an atomic variable...\n")
	increment := func(s string){
		for i:=0;i<5;i++{
			randomSleep()
			atomic.AddInt64(&aCount, 1) // Atomically increments aCount, blocking other processes from doing so
			fmt.Printf("%d - %s: Counter:%d\n",i+1, s,aCount)
		}
		wgA.Done()
	}

	wgA.Add(2)
	go increment("foo")
	go increment("bar")
	wgA.Wait()
}

func main() {
	// We have two additional go routines (if we ignore main()!)
	// so we add 2 to the wait group. Each wg.Done() execution drops
	// this count to zero, upon which the waitGroup exits, allowing main()
	// to terminate as well
	wg.Add(2)
	go foo() // Adding the 'go' keyword renders the subsequent func call concurrent
	go bar()
	wg.Wait() // The Wait() function waits for each function's .Done() to occur

	// Generate a race condition
	race_condition()

	// Eliminate race condition by using mutex
	useMutex()

	// Use an atomic variable instead (not recommended practice, use mutexs and channels instead
	atomicVariables()
}
