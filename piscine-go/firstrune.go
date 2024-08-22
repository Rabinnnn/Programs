package piscine

func FirstRune(s string) rune {
	if len(s) > 0 {
		// Convert the first character of the string to a rune
		return []rune(s)[0]
	}
	// Return a default value if the string is empty
	return 0
}
