package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printDigit(n int) {
	if n < 0 || n > 9 {
		return
	}
	digits := []rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	digit := digits[n]
	z01.PrintRune(digit)
}

func printValue(value int) {
	if value < 0 {
		z01.PrintRune('-')
		value = -value
	}

	if value < 10 {
		printDigit(value)
	} else {
		printValue(value / 10)
		printDigit(value % 10)
	}
}

func printString(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
}

func main() {
	points := &point{}

	setPoint(points)

	x, y := points.x, points.y

	printString("x = ")
	printValue(x)
	printString(", y = ")
	printValue(y)
	z01.PrintRune('\n')
}
