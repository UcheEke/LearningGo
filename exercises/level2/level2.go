package level2
import "fmt"

func divBy2(num int) (float64,bool)  {
	return float64(num) / 2.0 , num % 2 == 0
}

func max(data ...int) int {
	var result int
	for _, val := range data {
		if val > result {
			result = val
		}
	}
	return result
}

func Exercises(){
	fmt.Println("Level 2")
	// Exercise 1
	fmt.Println("Exercise 1")
	var num1, num2 int = 23,16
	result, evenOdd := divBy2(num1)
	fmt.Printf("%d divided by 2 is %.1f\n%d is even: %v\n", num1, result, num1, evenOdd)
	result, evenOdd = divBy2(num2)
	fmt.Printf("%d divided by 2 is %.1f\n%d is even: %v\n", num2, result, num2, evenOdd)

	// Exercise 2
	// Same as above using a func expression instead
	fmt.Println("\nExercise 2")
	divideIt := func(n int) (float64, bool) {
		return float64(n) / 2.0, n % 2 == 0
	}
	result, evenOdd = divideIt(num1)
	fmt.Printf("%d divided by 2 is %.1f\n%d is even: %v\n", num1, result, num1, evenOdd)
	result, evenOdd = divideIt(num2)
	fmt.Printf("%d divided by 2 is %.1f\n%d is even: %v\n", num2, result, num2, evenOdd)

	// Exercise 3 (and Exercise 5!)
	// Function that finds the maximum value in a list of integers
	fmt.Println("\nExercise 3")
	data := []int{2,5,4,7,56,8,43,20,1,16,74}
	fmt.Printf("The biggest value in %v is %d\n", data, max(data...))

	// Exercise 4
	fmt.Println("I know, I know, I know this much is", (true && false) || (false && true) || !(false && false)) // should print true

	// Exercise 6
	// Project Euler Rule #1: Thou shalt not divulgeth neither content nor solution of ProjectEuler.net
	// problems on any public forum, github.com in particular.
	//
	// Forsooth, this exercise doth unashamedly mock the creators of that goodly resource! :)
}
