package piscine

func JumpOver(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	var result string
	for i, char := range str {
		if (i+1)%3 == 0 {
			result += string(char)
		}
	}

	if len(result) == 0 {
		return "\n"
	}
	return result + "\n"
}
