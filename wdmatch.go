package main

import (
	"fmt"
	"os"
	
)

func main() {

	if len(os.Args) != 3{
		return
	}
	
 wdmatch(os.Args[1],os.Args[2])
}

func wdmatch(s1,s2 string){
	count := 0
	res := ""
	for i := 0; i < len(s1); i++{
		for j := count; j < len(s2); j++{
			if s1[i] == s2[j]{
				res += string(s1[i])
				count++
				break
			}
			count++
		}
	}
	if res == s1{	
		fmt.Println(s1)
	}
	return 
}