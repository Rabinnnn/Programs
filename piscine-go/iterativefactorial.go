package piscine

func IterativeFactorial(nb int) int {
	// Check for non-possible values
	if nb < 0 || nb >= 24 {
		return 0
	}
	result := 1
	for i := 1; i < nb+1; i++ {
		result = result * i
	}
	return result
}
