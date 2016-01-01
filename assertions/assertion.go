package main

// Go can assert whether a particular interface is of type T
// via the syntax
// interface.(T)
// This allows an interface to be processed as type T e.g

import "fmt"

// Create an empty interface type
type mytype interface{}

func main(){
	var mt mytype
	mt = 12

	fmt.Println(15 + mt.(int)) // This will pass

	// fmt.Println(15 + mt) // This will fail!
}
