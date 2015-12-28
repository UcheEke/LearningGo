package main

import (
	"fmt"
	"strings"
)

func hello(){
	fmt.Println("Hello, World!")
	fmt.Println()
}

func helloUche(){
	fmt.Println("Hello, my name is Uche")
	fmt.Println()
}

func helloYourName(){
	var name string
	fmt.Print("What is your name? ")
	fmt.Scanf("%s", &name)
	name = strings.TrimSpace(name)
	fmt.Printf("Hello %s\n", name)
	fmt.Println()
}

func allEvenNumbers(){
	for i := 1; i < 100; i += 2 {
		fmt.Printf("%d ",i)
	}
	fmt.Println()
}

func remainder() {
	var a,b int
	fmt.Print("Enter an integer > 1: ")
	fmt.Scanf("%d",&a)
	fmt.Print("Enter a non-zero integer smaller than the previous number: ")
	fmt.Scanf("%d", &b)

	fmt.Printf("\nThe remainder from %d/%d is %d\n", a,b,a%b)
	fmt.Println()
}

func fizzbuzz(){
	for i := 0; i <= 100; i++ {
		if (i % 3 == 0) && (i % 5 == 0) {
			fmt.Print("Fizzbuzz! ")
		} else if (i % 3 == 0){
			fmt.Print("Fizz ")
		} else if (i % 5 == 0) {
			fmt.Print("Buzz ")
		} else {
			fmt.Printf("%d ",i)
		}
	}
	fmt.Println()
}

func multipleSum() {
	var sum int
	var limit int = 1000
	for i := 0; i < limit; i++ {
		if (i % 3 == 0) || (i % 5 == 0) {
			sum += i
		}
	}

	fmt.Printf("The sum of all multiples of 3 or 5 below %d is %d\n", limit, sum)
	fmt.Println()
}

func main(){
	fmt.Println("Answers to exercises: ")
	// Exercise 1
	hello()
	// Exercise 2
	helloUche()
	// Exercise 3
	helloYourName()
	// Exercise 4
	remainder()
	// Exercise 5
	allEvenNumbers()
	// Exercise 6
	fizzbuzz()
	// Exercise 7
	multipleSum()
}