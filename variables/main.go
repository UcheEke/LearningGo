package main

import "fmt"

// Package level scope, x is a global variable
// In this case 'var x' assumes a type from the assignment
var x = "This is a string, defined using the var keyword in the global namespace"
// You could initialise the variable x in formal long form as:
// 'var x string = "initial string"'

type info struct {
    result string
}

func infoRepo(a int) (string, error) {
    return fmt.Sprintf("infoRepo return string and passed function variable a=%d\n",a), nil
}

func main() {
	fmt.Printf("%v \n", x) // uses the package level variable x

	// Short form variable declaration ':=' can only be used within functions
	a := 10         // Integer
	b := "golang"   // String
	c := 4.16       // Floating point number
	d := true       // boolean

    // We can initialise the struct 'info' and assign a value to a pointer to it like this:
    inf := info{}
    pInf := &inf

    pInf.result = "Assigned a value to info struct's 'result' parameter"

	// The print verb "%v" shows the value of the local variables a-d
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

    val, _ := infoRepo(2)
    fmt.Printf("%v", val)
    fmt.Println(pInf.result)

    // Casting
    // Go does not allow type numeric type mixing. Explicit casting is required if joint operations
    // on different types are required. Implicit promotion to the widest type doesn't occur
    // e.g. adding an int to a float
    var e int = 4
    var f float32 = 3.142

    fmt.Println("e = ", e, ", f = ", f, ", e + int(f) = ", e + int(f), ", float(e) + f = ",float32(e) + f)
}