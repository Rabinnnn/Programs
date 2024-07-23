package main

import (
	"fmt"
)

func main() {
	fmt.Println(AddFront("Hello", []string{"world"}))
	fmt.Println(AddFront("Hello", []string{"world", "!"}))
	fmt.Println(AddFront("Hello", []string{}))
}

func AddFront(s string, slice []string) []string {
	var result []string
	result = append(result, s)

	result = append(result, slice...)

	return result
}
