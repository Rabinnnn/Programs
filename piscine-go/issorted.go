package piscine

// IsSorted checks if the slice is sorted using the provided comparison function.
func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	descending := true

	for i := 0; i < len(a)-1; i++ {
		cmp := f(a[i], a[i+1])

		if cmp > 0 {
			ascending = false
		} else if cmp < 0 {
			descending = false
		}

		// If neither ascending nor descending, no need to continue checking
		if !ascending && !descending {
			return false
		}
	}

	return true
}

// Compare function for testing IsSorted
func f(a, b int) int {
	if a < b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}
