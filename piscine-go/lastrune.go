package piscine

func LastRune(s string) rune {
	if len(s) > 0 {
		// Convert the string to a slice of runes and return the last element
		runes := []rune(s)
		return runes[len(runes)-1]
	}
	// Return a default value if the string is empty
	return 0
}
