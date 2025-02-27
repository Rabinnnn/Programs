package main

import (
	"fmt"
)

func main() {
	fmt.Println(AlphaPosition('a'))
	fmt.Println(AlphaPosition('z'))
	fmt.Println(AlphaPosition('B'))
	fmt.Println(AlphaPosition('Z'))
	fmt.Println(AlphaPosition('0'))
	fmt.Println(AlphaPosition(' '))
}

func AlphaPosition(c rune) int {
	s1 := ""
	for i := 'a'; i <= 'z'; i++ {
		s1 += string(i)
	}

	s2 := ""
	for j := 'A'; j <= 'Z'; j++ {
		s2 += string(j)
	}

	output := -1
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if string(s1[i]) == string(c) || string(s2[i]) == string(c) {
			output = i + 1
		}
	}

	return output
}
