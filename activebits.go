package main

import (
	"fmt"
)

func main() {
	fmt.Println(ActiveBits(7))
}

func ActiveBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1 // this will increment count by 1 if n & 1 returns 1 or by 0 if n & 1 returns 0
		n >>= 1        // right shift to handle the next bit
	}
	return count
}
