package piscine

func isLatinLetter(c rune) bool {
	// Check if the character is a Latin letter
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func AlphaCount(s string) int {
	count := 0
	// Iterate through the string and count Latin letters
	for _, char := range s {
		if isLatinLetter(char) {
			count++
		}
	}
	return count
}
