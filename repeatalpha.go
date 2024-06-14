package main

import(
	"fmt"
	"os"
)
func main(){
	str := os.Args[1]
	// for _, char := range str{
	// 	fmt.Println(Indexx(char))

	// }
fmt.Println(Ralpha(str))
}

func Ralpha(str string)string{
	result := ""
	for _, char := range str{
		for i := 0; i < Indexx(char)+1; i++{
			result += string(char)
		}
	}
return result
}

func Indexx(r rune)int{
	var result string
	for i := 'a'; i <= 'z'; i++{
		result += string(i)
	}
	
	ind := 0

	for i := 0; i < len(result); i++{
		if string(result[i]) == string(r){
			ind = i
		}
	}
return ind
}
