package main

import "fmt"

func main() {
	//maps are defined with "map[key type]value type". There are a few ways of initialising:
	// 1. nil maps (since maps are reference types)
	var m1 map[string]int // this will be nil, to which you can assign nothing!
	// To my knowledge, there is no *common* usage of m1!

	// 2. empty maps
	var m2 = make(map[string]int) // creates a map of int values accessible via string keys
	// Unlike m1, m2 can have keys and values assigned to it
	m2["key1"] = 5
	m2["key2"] = 10

	// 3. map literal initialisation
	m3 := map[string]int{"key1" : 1, "key2" : 2}

	fmt.Println(m1,m2,m3)

	// Deleting keys from maps
	if _, ok := m2["key2"]; ok {
		delete(m2,"key2")
	} // Checks whether a particular key exists and deletes it
	fmt.Println("m2 with 'key2' deleted is", m2)
	// Note that the blank identifier would hold the value associated with the key if found
	// If the key is not present, the value would be the zero-value for the value type e.g.
	// m2 has int values so the following prints 0
	val, _ := m2["key2"]
	fmt.Println("Returned value for m2['key2'] is", val)

}
