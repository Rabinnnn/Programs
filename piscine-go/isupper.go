package piscine

func IsUpper(s string) bool {
	// Iterate through each character in the string
	for _, char := range s {
		// Check if the character is not an uppercase letter
		if char < 'A' || char > 'Z' {
			return false
		}
	}
	// All characters are uppercase
	return true
}
