package main
import "fmt"

// average takes in a variable number of float64 args, returning the simple mean of the
// input
func average(data ...float64) float64 {
	if len(data) > 0 {
		total := 0.0
		for _, val := range data {
			total += val
		}
		return total / float64(len(data))
	} else {
		return 0
	}
}

func main() {
	// Using a variadic function explicitly
	fmt.Printf("The average of the numbers [23,42,63,48,88,13,51]: %.2f\n", average(23,42,63,48,88,13,51))

	// Using the variadic functions implicitly
	data := []float64{23,42,63,48,88,13,51}
	fmt.Printf("The average of the numbers in data is %.2f\n", average(data...))
}
