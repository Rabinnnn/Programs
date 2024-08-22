package main

import (
	"fmt"
	"os"
	//"strings"
	//"strconv"
)

func main(){
	if len(os.Args) != 4{
		return
	}
	fmt.Println(srep(os.Args[1],os.Args[2],os.Args[3]))
}

func srep(s1 string, s2 string, s3 string)string{
	output := ""
	r1 := []rune(s1)
	r2 := []rune(s2)
	r3 := []rune(s3)

	for _, char := range r1{
		if char == r2[0]{
			char = r3[0]
		}
		output += string(char)
	}
	return output
}
