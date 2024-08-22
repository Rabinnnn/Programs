package piscine

func ToLower(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			result += string(char - 'A' + 'a')
		} else {
			result += string(char)
		}
	}
	return result
}
