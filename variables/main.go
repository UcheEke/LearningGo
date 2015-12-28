package main

import "fmt"

// Package level scope, x is a global variable
var x = "This is a string, defined using the var keyword in the global namespace"


func main() {
	fmt.Printf("%v \n", x) // uses the package level variable x

	// Assignments ':=' can only be used within functions
	a := 10
	b := "golang"
	c := 4.16
	d := true

	// The print verb "%v" shows the value of the local variables a-d
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}