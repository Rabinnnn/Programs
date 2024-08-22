package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb >= 24 {
		return 0 // Non-possible value
	}

	// Base case: factorial of 0 is 1
	if nb == 0 {
		return 1
	}

	// Recursive case: nb! = nb * (nb-1)!
	return nb * RecursiveFactorial(nb-1)
}
