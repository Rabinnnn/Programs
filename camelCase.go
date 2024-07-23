package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))

	fmt.Println(CamelToSnakeCase("132"))
	fmt.Println(CamelToSnakeCase("thisIsaTestOfCamelCase"))
}

func CamelToSnakeCase(s string) string {
	// str := ""
	// runes := []rune(s)
	if s == "" || !IsCamel(s) {
		return s
	}
	output := ""
	for i, char := range s {
		if i != 0 && i != len(s)-1 && !(char >= 'a' && char <= 'z') {
			output += "_" + string(char)
		} else if (char >= 'a' && char <= 'z') || i == 0 || i == len(s)-1 {
			output += string(char)
		}
	}
	return output
}

func IsCamel(s string) bool {
	// count := 0
	for i, char := range s {
		if i == len(s)-1 && !(char >= 'a' && char <= 'z') {
			return false
		} else if !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') {
			return false
		} else if !(s[i] >= 'a' && s[i] <= 'z') && !(s[i+1] >= 'a' && s[i+1] <= 'z') {
			return false
		}
	}
	return true
}
