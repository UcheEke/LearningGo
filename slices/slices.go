package main

import "fmt"

func main(){
	// Illustration of slice creation, expansion and capacity
	// Create an empty slice
	slice := make([]int, 0,5)
	fmt.Println("Initialised slice:",slice)

	// Appending beyond the length of the slice
	// causes the doubling of the capacity to accommodate the overflow
	// This has an effect on the performance since the underlying array
	// the slice references has to be exchanged for a larger one
	for i := 0; i < 10; i++ {
		slice = append(slice,i)
		fmt.Printf("%v Len = %v Capacity = %v\n",i, len(slice), cap(slice))
	}

	// Differences in initialisation
	// Using 'var' to initialise a slice results in a nil value e.g.
	var nilSlice []string
	fmt.Printf("It's %v that nilSlice is nil!\n", nilSlice == nil)

	// Initialising with empty braces results in an empty slice with 0 capacity
	// e.g
	var emptySlice = []string{}
	fmt.Printf("Saying emptySlice is nil is patently %v\n", emptySlice == nil)
	fmt.Printf("emptySlice has a length %v and capacity of %v\n", len(emptySlice), cap(emptySlice))
	// and in order to add elements to this slice, you have to append() to it. No indexing is
	// possible beforehand

	// A slice can also be initialised from an array initialisation using the new() function e.g.
	var arraySlice = new([50]int)[0:10]
	fmt.Printf("arraySlice is a []int of length %v, capacity %v as shown here: %v\n", len(arraySlice), cap(arraySlice), arraySlice)

	// changing the size of a slice can be done via the append function and selective ranges
	mySlice := make([]int, 10)
	for i:=0; i < len(mySlice); i++ {
		mySlice[i] = i
		mySlice[i]++ // Increment each value within the slice
	}

	fmt.Printf("mySlice: %v\n", mySlice)
	// I want to remove the numbers 5-7 from the array so I can do the following
	mySlice = append(mySlice[0:4], mySlice[7:]...) // Note the ellipsis effectively expanding mySlice[7:] into elements
	fmt.Printf("mySlice (after edit): %v\n",mySlice)

}
