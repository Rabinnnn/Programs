package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	// Create an array to count occurrences of each digit
	count := [10]int{}

	// Count occurrences of each digit
	for n != 0 {
		count[n%10]++
		n /= 10
	}

	// Print digits in ascending order
	for i := 0; i <= 9; i++ {
		for count[i] > 0 {
			z01.PrintRune(rune(i) + '0')
			count[i]--
		}
	}
}
