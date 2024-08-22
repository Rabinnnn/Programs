package piscine

// PodiumPosition takes a slice of slices of type string and returns the competitor positions correctly.
func PodiumPosition(podium [][]string) [][]string {
	// Initialize a map to store the correct positions
	positionMap := map[string]string{
		"1st": "1st Place",
		"2nd": "2nd Place",
		"3rd": "3rd Place",
		"4th": "4th Place",
	}

	// Initialize a 2D array with the correct number of competitors
	result := [][]string{
		{positionMap["1st"]},
		{positionMap["2nd"]},
		{positionMap["3rd"]},
		{positionMap["4th"]},
	}

	// Iterate over the original podium positions
	for i := 0; i < len(podium); i++ {
		// Check if the current position is valid
		if len(podium[i]) > 0 {
			// Extract the position (1st, 2nd, 3rd, etc.) from the competitor's position slice
			position := podium[i][0]

			// Look up the correct position from the map and store it in the result slice
			result[i][0] = positionMap[position]
		}
	}

	return result
}
