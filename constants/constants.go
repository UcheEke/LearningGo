package main

import "fmt"


// individual declarations
const p = "death and taxes"
const q int = 42

// Multiple declarations
const (
	Pi = 3.14
	Language = "Go"
)

// iotas
const (
	A	= iota	// 0
	B			// 1
	C			// 2
)

// Redefining an iota in another const block resets the iota to 0
const (
	D	= iota	// 0
	E			// 1
	F			// 2
)

// Define with expressions
const (
	_ = iota // 0
	KB = 1 << (iota * 10) // = 1024 = 2**10
	MB = 1 << (iota * 10) // 2**20
	GB = 1 << (iota * 10) // 2**30
)


func main() {
	fmt.Printf("p (type %T) = %q\n", p,p)
	fmt.Printf("q (type %T) = %d\n", q,q)
}
