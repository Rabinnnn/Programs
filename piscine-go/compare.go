package piscine

func Compare(a, b string) int {
	// Compare the two strings lexicographically
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	// If one string is a prefix of the other, return the difference in lengths
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}

	// The strings are equal
	return 0
}
