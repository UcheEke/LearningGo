package main

import (
	"fmt"
)

func do() {
	var err error
	defer func() {
		if err != nil {
			fmt.Printf("Something went wrong: %s\n", err)
		}
	}()

	err = fmt.Errorf("This is something going wrong...")
}

func main() {
	do()
}
