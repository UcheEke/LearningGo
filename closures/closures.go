package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func makeGreeter(greeting string) func(name string) {
	return func(name string) {
		fmt.Printf("%s %s!\n", greeting, name)
	}
}

func main(){
	// Simple closure example
	increment := wrapper()
	fmt.Println(increment()) // Prints 1
	fmt.Println(increment()) // Prints 2

	// Example of a slightly more complex closure
	greet := makeGreeter("Oh, Hello")
	greet("Uche")
	greet("Barbara")
	greet("Thomas")
}
