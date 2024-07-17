package main

import (
	"fmt"
	"os"
	"strconv"
)
func main(){
	if len(os.Args) != 2 {
		
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 1{
		
		return
	}
	prime := ""

	for i := 2; i <= num; i++{
		for Isp(i) && num%i == 0{
			prime += strconv.Itoa(i) + "*" 
			num = num/i
		}
	}
	for j := 0; j < len(prime)-1; j++{
		fmt.Print(string(prime[j]))
	}
	
}

func Isp(n int)bool{
	
	for i := 2; i*i <= n; i++{
		if n%i == 0{
		 return	false
		}
	}
	return true
}