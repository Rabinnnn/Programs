package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Command line arguments
	args := os.Args[1:]

	// Iterate through ASCII values and print matching arguments
	for i := '0'; i <= 'z'; i++ {
		for _, arg := range args {
			if arg[0] == byte(i) {
				z01.PrintRune('\n')
				for _, char := range arg {
					z01.PrintRune(char)
				}
			}
		}
	}
	z01.PrintRune('\n')
}
