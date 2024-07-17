package main

import (
	"fmt"
	//"sort"
	"os"
	// "sort"
	//"strconv"
)
func main(){

	if len(os.Args) != 3{
		return
	}
	fmt.Println( Inter(os.Args[1], os.Args[2]) )
	
}

func Inter(s1,s2 string)string{
	var result string
	seen := make(map[byte]bool)

	present := make(map[byte]bool)
	for i := 0; i < len(s2); i++{
		present[s2[i]] = true
	}

	for j := 0; j < len(s1); j++{
		if present[s1[j]] && !seen[s1[j]]{
			result += string(s1[j])
			seen[s1[j]] = true
		}
	}
return	result
}

