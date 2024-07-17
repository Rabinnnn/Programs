package main

import (
	"fmt"
	//"strconv"
	//"sort"
	"os"
	// "sort"
	//"strconv"
)

func main() {

	if os.Args[1] == ""{
		fmt.Println("1")
		return
	}
fmt.Println(wdmatch(os.Args[1],os.Args[2]))
}

func wdmatch(s1,s2 string)int{
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
	fmt.Println(res)
	if res == s1{
		return 1
	}
	return 0
}