package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	guess := 1
	for guess*guess <= nb {
		if guess*guess == nb {
			return guess
		}
		guess++
	}

	return 0
}
