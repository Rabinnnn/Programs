package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		// Read from standard input and print to standard output
		copyAndPrint(os.Stdin, os.Stdout)
		return
	}

	// Process each file specified as an argument
	for _, fileName := range args {
		file, err := os.Open(fileName)
		if err != nil {
			printString("ERROR: " + err.Error() + "\n")
			os.Exit(1) // Exit with a non-zero status
		}
		defer file.Close()

		// Read from the file and print to standard output
		copyAndPrint(file, os.Stdout)
	}
}

func copyAndPrint(src io.Reader, dest io.Writer) {
	_, err := io.Copy(dest, src)
	if err != nil {
		printString("ERROR: " + err.Error() + "\n")
		os.Exit(1) // Exit with a non-zero status
	}
}

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
