package main

import "fmt"

const metersToYards float64 = 1.09361

func main(){
	a := 43
	fmt.Println("a's value:", a)
	fmt.Printf("a's address (hex): %#x\n", &a)
	fmt.Printf("a's address in decimal: %d\n", &a)

	// Using addresses: Best illustrated simply by taking user input and assigning
	// it to the address of a predefined variable e.g.

	var meters float64
	fmt.Print("Enter a distance (m): ")
	fmt.Scanf("%f", &meters)

	fmt.Printf("Distance %.2fm = %.2f'\n", meters, meters * metersToYards)
}


