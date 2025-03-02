package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	for _, arg := range os.Args[1:]{
	
		if isBracketed(arg){
			fmt.Println("OK")
		}else{
			fmt.Println("ERROR")
		}
	}
}

func isBracketed(str string) bool {
	var openBrackets []string
	pointer := -1
	count := 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "(" || string(str[i]) == "[" || string(str[i]) == "{" {
			count++
			openBrackets = append(openBrackets, string(str[i]))
			pointer++
		} else {
			if string(str[i]) == ")" && openBrackets[pointer] == "(" {
				count++
				openBrackets = openBrackets[:pointer]
				pointer--
			}else if string(str[i]) == "]" && openBrackets[pointer] == "[" {
				count++
				openBrackets = openBrackets[:pointer]
				pointer--
			}else if string(str[i]) == "}" && openBrackets[pointer] == "{" {
				count++
				openBrackets = openBrackets[:pointer]
				pointer--
			}
		}
	}
	if count == 0{
		return true
	}

	return len(openBrackets) == 0
}
