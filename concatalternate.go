package main

import (
	"fmt"
)

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}

func ConcatAlternate(slice1, slice2 []int) []int {
	var output []int
	firstSlice, secondSlice := First(slice1, slice2)
	count := 0
	for i := 0; i < len(firstSlice); i++ {
		output = append(output, firstSlice[i])
		for j := count; j < len(secondSlice); j++ {
			output = append(output, secondSlice[j])
			count++
			break
		}
	}
	return output
}

func First(slice1, slice2 []int) ([]int, []int) {
	if len(slice2) > len(slice1) {
		return slice2, slice1
	} else if len(slice1) > len(slice2) {
		return slice1, slice2
	}
	return slice1, slice2
}
