package main

import "fmt"

func zero(x int) {
	fmt.Printf("zero: Address of pointer = %p\n", &x) // verb %p ('p' for pointer) is similar to %#x (prints base 16 number with 0x prefix)
	x = 0
}

func zeroPoint(x *int) {  // <- takes this pointer
	fmt.Printf("zeroPoint: Address of pointer = %p\n", x)
	*x = 0 // assignment via dereference
}

func main() {
	a := 43
	fmt.Println(a) 	// 43
	fmt.Println(&a) // -> addr(a)
	var b *int = &a  // *int is an integer pointer
	fmt.Println(b)	// -> addr(a)
	fmt.Println(*b) // 43 ('*' in this case dereferences the address, providing the value)

	*b = 42	// i.e. assign 42 to the address b is pointing to

	zero(a)
	fmt.Println("Using zero(): a =",a) // a is still 43
	zeroPoint(&a)
	fmt.Println("Using zeroPoint(): a =",a) // a is now 0
}