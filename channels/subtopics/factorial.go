package subtopics

import _"fmt"

func Factorial(num int) int {

	factCh := make(chan int)

	go func(){
		product := 1
		for i:=num; i > 0; i-- {
			product *= i
		}
		factCh <- product
	}()

	return <- factCh
}
