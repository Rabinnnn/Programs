package main

import (
	"flag"
	"fmt"
	"strings"
)

// ANSI color codes
var colors = map[string]string{
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"magenta": "\033[35m",
	"cyan":   "\033[36m",
	"reset":  "\033[0m",
}

// Main function
func main() {
	// Define flags
	colorFlag := flag.String("color", "", "Color to use for highlighting (e.g., red, green, blue, etc.)")
	flag.Parse()

	// Check if the required flags and arguments are provided
	if len(flag.Args()) < 2 {
		printUsage()
		return
	}

	// Extract the substring and text
	substring := flag.Arg(0)
	text := flag.Arg(1)

	// Validate color
	colorCode, exists := colors[*colorFlag]
	if !exists {
		printUsage()
		return
	}

	// Color the substring in the text
	coloredText := highlightSubstring(text, substring, colorCode)
	fmt.Println(coloredText)
}

// Usage message
func printUsage() {
	fmt.Println("Usage: go run . --color=<color> <substring> <string>")
	fmt.Println("Example: go run . --color=red kit \"a king kitten have kit\"")
}

// Function to highlight the substring with the specified color
func highlightSubstring(text, substring, colorCode string) string {
	var result strings.Builder
	index := 0

	for {
		start := strings.Index(text[index:], substring)
		if start == -1 {
			result.WriteString(text[index:])
			break
		}

		start += index
		result.WriteString(text[index:start])
		result.WriteString(colorCode + substring + colors["reset"])
		index = start + len(substring)
	}

	return result.String()
}
