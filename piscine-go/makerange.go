package piscine

func MakeRange(min, max int) []int {
	// Check if min is greater than or equal to max
	if min >= max {
		return nil
	}

	// Calculate the size of the slice
	size := max - min

	// Create the slice with the specified size
	result := make([]int, size)

	// Fill the slice with values from min to max-1
	for i := 0; i < size; i++ {
		result[i] = min + i
	}

	return result
}
