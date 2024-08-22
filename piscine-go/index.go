package piscine

func Index(s string, toFind string) int {
	// Iterate through the string s
	for i := 0; i <= len(s)-len(toFind); i++ {
		// Check if toFind matches a substring of s starting at index i
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	// If toFind is not found, return -1
	return -1
}
