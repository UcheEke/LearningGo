package main

import (
	"fmt"
	"time"
)

// An error is anything that implements the Error interface
// which requires it have an "Error() string {}" method
type MyError struct {
	str string
	when time.Time
}

func (m MyError) Error() string {
	return fmt.Sprintf("%s at %s\n", m.str, m.when)
}

func main(){
	err := MyError{}
	err.str = "Something when wrong"
	err.when = time.Now()

	fmt.Printf(err.Error())
}


