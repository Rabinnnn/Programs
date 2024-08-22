package piscine

func IsAlpha(s string) bool {
	// Iterate through each character in the string
	for _, char := range s {
		// Check if the character is not an alphanumeric character
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return false
		}
	}
	// All characters are alphanumeric or the string is empty
	return true
}
