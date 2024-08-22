package piscine

func BasicAtoi2(s string) int {
	result := 0

	for _, char := range s {
		// Check if the character is a digit
		if char >= '0' && char <= '9' {
			digit := int(char - '0')
			result = result*10 + digit
		} else {
			// If a non-digit character is encountered, return 0
			return 0
		}
	}

	return result
}
