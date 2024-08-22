package piscine

func Split(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	var result []string
	substring := ""
	sepLength := len(sep)

	for i := 0; i < len(s); i++ {
		if i+sepLength <= len(s) && s[i:i+sepLength] == sep {
			result = append(result, substring)
			substring = ""
			i += sepLength - 1
		} else {
			substring += string(s[i])
		}
	}

	result = append(result, substring)

	return result
}
