package main

import "fmt"

func main() {
	// The difference b/w an array and a slice in declaration is arrays
	// have a size within the square brackets
	var x [20]int

	for i := 0; i < len(x); i++ {
		x[i] = i * i
	}
	fmt.Println(x)

	var y [58]string

	for i := 65; i <= 122; i++ {
		y[i-65] = string(i)
	}

	fmt.Println(y)
}
