package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main(){
	increment := wrapper()
	fmt.Println(increment()) // Prints 1
	fmt.Println(increment()) // Prints 2
}
