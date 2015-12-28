package main

import "fmt"

func main(){
	// Conversion of string into a slice of runes
	phrase := "Hello, World"
	charList := []byte(phrase)
	fmt.Println("'Hello,World' in bytes:",charList)

	// Listing utf-8 characters
	// Note how non-ascii characters use more than 1 byte (upto 3 more)
	// a rune is an alias for int32 or a 4-byte character representation
	for i := 0; i < 400; i++ {
		fmt.Printf("%d - %s - %v\n",i,string(i),[]byte(string(i)))
	}
}
