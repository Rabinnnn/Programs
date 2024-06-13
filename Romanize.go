package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main(){
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num >= 4000{
		fmt.Println("Error: cannot convert to roman digit")
		return
	}
	fmt.Println(ToRoman(num))
}

func ToRoman(num int)string{
	roman := map[int]string{
		1: "I",
		4: "IV",
		5: "V",
		9 : "IX",
		10: "X",
		40: "XL",
		50: "L",
		90: "XC",
		100: "C",
		400: "CD",
		500: "D",
		900: "CM",
		1000: "M",
	}

	//create a slice of integers and populate it with key values
	var keys []int
	for key := range roman{
		keys = append(keys, key)
	}

	//sort the keys in ascending order
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	result := ""
	calc := ""

	//convert to roman based on if and how many times a specific key goes into the num
	for _, key := range keys{
		value := roman[key]
		remainder := num%key

		for i := 0; i < num/key; i++{
			result += value	
			//show how the calculation is done		
			if len(value) == 1{
				if calc == ""{
					calc = calc + value 
				}else{
					calc = calc + "+" + value
				}
			}else if len(value) == 2{
				if calc == ""{
					calc = calc + "(" + string(value[1]) + "-" + string(value[0]) + ")"
				}else{
					calc = calc + "+" + "(" + string(value[1]) + "-" + string(value[0]) + ")"
				}
			}
			
		}
		num = remainder
		
		//break the loop when the conversion is over
		if remainder == 0{
			break
		}

	}
	//print calculation and return result too
	fmt.Println(calc)
	return result
}