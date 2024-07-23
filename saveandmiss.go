package main

import (
	"fmt"
	// "strings"
	//"strconv"
)

func SaveAndMiss(arg string, num int) string {
	output := ""
	chunk := ""
	if num < 1 {
		return arg
	}
	//count := 0
	for i := 0; i < len(arg); i++ {
		chunk += string(arg[i])
		if len(chunk) == num {
			output += chunk
			chunk = ""
			i = i + num
			//break
		}

	}
	if chunk != "" {
		output += chunk
	}
	return output
}

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}
