package main

import (
	"fmt"
)

func isPositive(num int) bool {
	return num > 0
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// Test isPositive function
	fmt.Println("isPositive(5):", isPositive(5))   // true
	fmt.Println("isPositive(-5):", isPositive(-5)) // false
	fmt.Println("isPositive(0):", isPositive(0))   // false

	// Test abs function
	fmt.Println("abs(5):", abs(5))       // 5
	fmt.Println("abs(-5):", abs(-5))     // 5
	fmt.Println("abs(0):", abs(0))       // 0
	fmt.Println("abs(-100):", abs(-100)) // 100
}
