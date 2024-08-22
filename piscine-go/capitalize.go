package piscine

func Capitalize(s string) string {
	result := ""
	upperNext := true

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0') && (char <= '9') {
			if upperNext {
				if char >= 'a' && char <= 'z' {
					result += string(char - 'a' + 'A')
				} else {
					result += string(char)
				}
				upperNext = false
			} else {
				if char >= 'A' && char <= 'Z' {
					result += string(char - 'A' + 'a')
				} else {
					result += string(char)
				}
			}
		} else {
			result += string(char)
			upperNext = true
		}
	}
	return result
}
