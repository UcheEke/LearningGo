package main

import "fmt"

func square(n int) int {
	return n * n
}

func cube(n int) int {
	return n * n * n
}

// Use a closure to generate various "Multiply by N" functions
func multiplyBy(n int) func(int) int {
	return func(val int) int {
		return n * val
	}
}

// dataProcessor takes an int slice and applies a valid callback function to each element
// valid => has the required function signature
func dataProcessor(data []int, callback func(int) int ) []int {
	var result []int
	for _, val := range data {
		result = append(result, callback(val))
	}
	return result
}

// dataFilter uses a filtering function to filter the data, returning a new slice with the elements in data
// that satisfied the filter test
func dataFilter(data []int, callback func(int) bool) []int {
	var result []int
	for _, val := range data {
		if callback(val) {
			result = append(result, val)
		}
	}
	return result
}

func main(){
	data := []int{1,2,3,4,5,6,7,8,9}

	// Generate a few "xN" functions
	x2 := multiplyBy(2)
	x3 := multiplyBy(3)
	x7 := multiplyBy(7)

	// apply the various functions to the data set
	fmt.Println("Initial data:", data)
	fmt.Println("Each data point squared:", dataProcessor(data,square))
	fmt.Println("Each data point cubed:", dataProcessor(data,cube))
	fmt.Println("Each data point x 2:", dataProcessor(data, x2))
	fmt.Println("Each data point x 3:", dataProcessor(data, x3))
	fmt.Println("Each data point x 7:", dataProcessor(data, x7))

	// Instead of using a closure, we'll define the filter functions inline...
	fmt.Println("Filter data for all x <= 4:", dataFilter(data, func(n int) bool {return n <= 4}))
	fmt.Println("Filter data for all x > 5:", dataFilter(data, func(n int) bool {return n > 5}))
}
