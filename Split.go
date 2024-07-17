package main

import (
	"fmt"
	//"strconv"
	//"sort"
	//"os"
	// "sort"
	//"strconv"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
func Split(s, sep string) []string {
	var res []string
	var word string
	for i := 0; i < len(s); i++ {
		if i+len(sep) < len(s) && s[i:i+len(sep)] == sep {
			res = append(res, word)
			word = ""
			i = i + len(sep) - 1
		} else {
			word += string(s[i])
		}
	}
	res = append(res, word)
	return res
}
