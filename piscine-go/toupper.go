package piscine

func ToUpper(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			// Convert lowercase letter to uppercase
			result += string(char - 'a' + 'A')
		} else {
			// Keep non-alphabetic characters unchanged
			result += string(char)
		}
	}
	return result
}
