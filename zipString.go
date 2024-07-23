package main

import (
	"fmt"
	"strings"
	//"strconv"
)

func ZipString(s string) string {
	words := strings.Fields(s)
	var result strings.Builder

	for i, word := range words {
		if i > 0 {
			result.WriteString("1 ")
		}
		result.WriteString(processWord(word))
	}

	return result.String()
}

func processWord(word string) string {
	var compressed strings.Builder
	charCount := make(map[rune]int)
	seen := make(map[rune]bool)

	// Count occurrences of each character in the word
	for _, char := range word {
		charCount[char]++
	}

	// Build the compressed word
	for _, char := range word {
		if !seen[char] {
			seen[char] = true
			compressed.WriteString(fmt.Sprintf("%d%c", charCount[char], char))
		}
	}

	return compressed.String()
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))

	fmt.Println(ZipString("aaaa"))
	fmt.Println(ZipString("zzzzzZZZZZZ"))
	fmt.Println(ZipString(""))
	fmt.Println(ZipString("ziiiiipMeee"))
	fmt.Println(ZipString("hhellloTthereYouuunggFelllas"))

}
