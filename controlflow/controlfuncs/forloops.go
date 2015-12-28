package controlfuncs

import "fmt"

func Fors() {
	// Basic for loop
	fmt.Println("Basic forloop 'for i := 0; i < limit; i++ {}'...")
	limit := 20
	for i := 0; i < limit; i++ {
		fmt.Println("i =", i)
	}

	// for with a condition === Clang while
	fmt.Println("Basic while type forloop 'for condition {}'...")
	fmt.Println("Decrementing limit from 20...")
	for limit > 0 {
		fmt.Printf("%d ",limit)
		limit--
	}
	fmt.Println()

	// Endless loop with for
	fmt.Println("Endless loop: press <enter> to break... ")
	var endkey rune // Initialise two runes
	for {
		fmt.Scanf("%c", &endkey)
		if endkey != '\x0D' { // returns are interpreted as spaces with Scanf
			fmt.Println("Exiting...")
			break
		}
	}

	// for with continue
	fmt.Println("Illustration of for with continue...")

	var i int
	for {
		i++
		if i % 2 == 0 {
			continue   // skips the even numbers (illustration only; i+2 with initial i=-1 would be simpler!)
		} else {
			if i >= 50 {
				break
			}
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}