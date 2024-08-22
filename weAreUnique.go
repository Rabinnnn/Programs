package main

import (
	"fmt"
	"strings"
)

func WeAreUnique(str1, str2 string) int {
	// Return -1 if both strings are empty
	if str1 == "" && str2 == "" {
		return -1
	}

	// Use maps to track unique characters in each string
	uniqueInStr1 := make(map[rune]bool)
	uniqueInStr2 := make(map[rune]bool)

	// Populate the map for the first string
	for _, ch := range str1 {
		uniqueInStr1[ch] = true
	}

	// Populate the map for the second string
	for _, ch := range str2 {
		uniqueInStr2[ch] = true
	}

	// Count characters that are unique to each string
	count := 0
	for ch := range uniqueInStr1 {
		if !strings.ContainsRune(str2, ch) {
			count++
		}
	}
	for ch := range uniqueInStr2 {
		if !strings.ContainsRune(str1, ch) {
			count++
		}
	}

	// Return 0 if there are no unique characters
	if count == 0 {
		return 0
	}

	return count
}

func main() {
	// Test cases
	fmt.Println(WeAreUnique("abc", "bcd"))    // Should return 2 ('a' and 'd')
	fmt.Println(WeAreUnique("hello", "world")) // Should return 7 ('h', 'e', 'w', 'r', 'd')
	fmt.Println(WeAreUnique("same", "same"))   // Should return 0
	fmt.Println(WeAreUnique("", ""))           // Should return -1
	fmt.Println(WeAreUnique("abc", ""))        // Should return 3 ('a', 'b', 'c')
	fmt.Println(WeAreUnique("", "abc"))        // Should return 3 ('a', 'b', 'c')
}
