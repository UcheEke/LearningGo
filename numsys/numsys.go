package main

import "fmt"

func decimal(number int) {
	// uses the '%d' format specifier for decimals
	fmt.Printf("%d\n", number)
}

func binary(number int) {
	// uses the '%b' format specifier for binary
	fmt.Printf("%d\t%b\n", number, number)
}

func hex(number int) {
	// %x is the basic format specifier for hex numbers
	// note that #x and #X format specifiers provide the '0x' and '0X' respoectively
	fmt.Printf("%d\t%b\t%#X\n", number, number, number)
}

func utfs(number int){
	// %q format specifier 'quotes' the utf-8 rune for a given integer in range
	fmt.Printf("%d \t %b \t %#x \t %q \n", number, number, number, number)
}

func main(){

	decimal(42)
	binary(42)
	hex(42)

	// Print out all the ascii characters from the utf-8 charset
	for i := 0; i < 128; i++ {
		utfs(i)
	}
}
