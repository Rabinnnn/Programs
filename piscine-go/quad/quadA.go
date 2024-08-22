package main

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 || i == y-1 {
				if j == 0 || j == x-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			} else {
				if j == 0 || j == x-1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
	//fmt.Printf("[quadA] [%d] [%d]\n", x, y)
}
func main() {
	// Use default dimensions for demonstration
	QuadA(3, 3)
}
