package main

import (
	"regexp"
	"fmt"
)

var mustRegEx = regexp.MustCompile("[0-9]+")

func main() {
	// MustCompile will replace this section of code below
	/*
	if err != nil {
		log.Fatalf("Bad regexp: %s\n", err)
	}*/

	fmt.Printf("%t\n", mustRegEx.MatchString("Can you point me to NGC1502A?"))
}
