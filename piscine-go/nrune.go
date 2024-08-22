package piscine

func NRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		// Return 0 if the index is out of bounds or negative
		return 0
	}
	// Convert the string to a slice of runes and return the nth element
	runes := []rune(s)
	return runes[n-1]
}
