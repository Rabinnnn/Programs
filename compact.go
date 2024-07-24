package main

import (
	"fmt"
)

const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}

func Compact(s *[]string) int {
	// dereference the pointr to get the actual slice
	slice := *s
	count := 0

	for i, elem := range slice {
		if elem != "" {
			slice[count] = slice[i]
			count++
		}
	}
	// truncate the slice to remove elements with zero value
	*s = slice[:count]
	return count
}
