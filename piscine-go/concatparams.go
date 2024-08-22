package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	// Calculate the total length of the concatenated string
	totalLength := 0
	for _, arg := range args {
		totalLength += len(arg)
	}

	// Create a slice to store the concatenated string
	result := make([]byte, totalLength+len(args)-1) // len(args)-1 for newline characters

	// Copy each argument and newline to the result slice
	offset := 0
	for i, arg := range args {
		copy(result[offset:], arg)
		offset += len(arg)
		if i < len(args)-1 {
			result[offset] = '\n'
			offset++
		}
	}

	return string(result)
}
