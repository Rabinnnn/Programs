package piscine

// Compact deletes elements with zero-values in the slice and returns the number of non-zero elements.
func Compact(ptr *[]string) int {
	original := *ptr
	nonZeroCount := 0

	// Iterate through the slice and compact (remove zero-values)
	j := 0
	for _, val := range original {
		if val != "" {
			original[j] = val
			j++
			nonZeroCount++
		}
	}

	// Update the length of the slice
	*ptr = original[:j]

	return nonZeroCount
}
