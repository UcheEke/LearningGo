package stringutil

// Reverse uses the unexported function 'reverse2' to reverse the
// characters in a string
func Reverse(s string) string {
	return reverse2(s)
}

func ReverseAgain(s string) string {
	return reverse3(s)
}
