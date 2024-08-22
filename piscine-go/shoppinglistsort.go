package piscine

// ShoppingListSort sorts the slice of strings based on string length.
// Strings within the input slice must be of different lengths.
func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if len(slice[i]) > len(slice[j]) {
				// Swap the elements
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}
