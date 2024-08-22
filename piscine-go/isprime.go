package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	// Check for divisibility from 2 to the square root of nb
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}

	return true
}
