package piscine

func IsLower(s string) bool {
	// Iterate through each character in the string
	for _, char := range s {
		// Check if the character is not a lowercase letter
		if char < 'a' || char > 'z' {
			return false
		}
	}
	// All characters are lowercase
	return true
}
