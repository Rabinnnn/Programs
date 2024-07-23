package main

import (
	"fmt"
)

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))

	input4 := []uint{5}
	fmt.Println(CanJump(input4))
}

func CanJump(a []uint) bool {
	if len(a) == 0 {
		return false
	} else if len(a) == 1 {
		return true
	}

	var nextIndex uint

	for i := 0; i < len(a); i++ {
		nextIndex = uint(i) + a[i]
		i = i + int(nextIndex)
		if int(nextIndex) > len(a)-1 {
			return false
		}
	}
	return true
}
