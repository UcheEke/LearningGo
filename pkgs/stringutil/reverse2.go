package stringutil

// reverse2 is only visible within the package stringutil.

// reverse2 takes in a string and returns a copy of the string with the
// character order reversed. This is more efficient than reverse3 since it changes the
// input string in place, using less memory
func reverse3(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < len(chars)/ 2; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i] // Swap the chars up to the midpoint of the string
	}
	return string(chars)
}

// reverse2 uses a less efficient method to reverse the characters in s cf reverse3
// by the end of this method, there are two slices of len(s) in memory
func reverse2(s string) string {
	var chars []rune
	str := []rune(s)
	for i := len(str) - 1; i >= 0; i-- {
		chars = append(chars, str[i])
	}
	return string(chars)
}