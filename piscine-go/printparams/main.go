package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Iterate through command line arguments starting from index 1
	// (index 0 is the program name)
	for _, arg := range os.Args[1:] {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
