package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	number, err := strconv.Atoi(os.Args[1])
	if err != nil || number < 0 {
		fmt.Println("ERROR")
		return
	}

	fmt.Println(intToHex(number))
}

func intToHex(n int) string {
	if n == 0 {
		return "0"
	}

	hexChars := "0123456789abcdef"
	result := ""
	for n > 0 {
		result = string(hexChars[n%16]) + result
		n /= 16
	}
	return result
}
