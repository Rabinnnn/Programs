package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	foundDigit := false

	for _, char := range s {
		if char >= '0' && char <= '9' {
			// Convert the digit character to an integer
			digit := int(char - '0')

			// Update the result by multiplying it by 10 and adding the new digit
			result = result*10 + digit

			foundDigit = true
		} else if char == '-' && !foundDigit {
			// If '-' is encountered before any digit, set the sign to negative
			sign = -1
		} else if char == '+' && !foundDigit {
			// If '+' is encountered before any digit, set the sign to positive
			sign = 1
		}
	}

	// Apply the sign to the result
	result *= sign

	return result
}
