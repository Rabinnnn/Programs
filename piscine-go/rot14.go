package piscine

func Rot14(s string) string {
	result := []rune(s)

	for i, char := range result {
		if 'a' <= char && char <= 'z' {
			result[i] = 'a' + (char-'a'+14)%26
		} else if 'A' <= char && char <= 'Z' {
			result[i] = 'A' + (char-'A'+14)%26
		}
	}

	for j, char := range result {
		if 'a' <= char && char <= 'z' {
			result[i] = 'a' + (char-'a'+14)%26
		} else if 'A' <= char && char <= 'Z' {
			result[i] = 'A' + (char-'A'+14)%26
		}
	}

	return string(result)
}
