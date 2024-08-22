package piscine

func IsNumeric(s string) bool {
	// Iterate through each character in the string
	for _, char := range s {
		// Check if the character is not a numerical digit
		if char < '0' || char > '9' {
			return false
		}
	}
	// All characters are numerical digits
	return true
}
