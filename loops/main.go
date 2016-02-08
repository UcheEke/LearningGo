package main

import "fmt"

func main(){
	var (
		numbers = []int{2,4,6,8,10,22,1,3,0,5,7,9,11}
		oddCount int
		evenCount int
	)

	exitLoop:
	for _, n := range numbers {
		switch {
		case n == 0:
			fmt.Println("Found a zero in the list!")
			break exitLoop
		case n % 2 == 0:
			evenCount++
		default:
			oddCount++
		}
	}

	fmt.Printf("Even numbers: %d\nOdd Numbers: %d\n", evenCount,oddCount)
}
