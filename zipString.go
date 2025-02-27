package main
import (
	"fmt"
	"strconv"
)


func ZipString(s string) string {
	words := Split(s)
	charCount := make(map[byte]int)
	count := 1
	var res string
	for j, word := range words{
		for i := 0; i < len(word); i++{
			if i != len(word)-1 && word[i+1] == word[i]{
				count++
			}else{
				charCount[word[i]] = count
				count = 1
			}
		}

		str := Rdup(word)
		for i := 0; i < len(str); i++{
			res += strconv.Itoa(charCount[str[i]]) + string(str[i])
		}
		if j != len(words)-1{
			res = res + "1 "
		}

		
		
	}
	
	
return res
}

func Rdup(s string)string{
	//var seen bool
	var res string
	for i := 0; i < len(s); i++{
		if i != len(s)-1 && s[i+1] == s[i]{
			continue
		}else{
			res += string(s[i])
		}

	}
	return res
}


func Split(s string)[]string{
	var res []string
	var word string
	for _, char := range s{
		if char == ' '{
			if word != ""{
				res = append(res, word)
				word = ""
			}
		}else {
			word += string(char)
		}
	}
	if word != ""{
		res = append(res, word)
		word = ""
	}
	return res
}
func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))

	fmt.Println(ZipString("aaaa"))
	fmt.Println(ZipString("zzzzzZZZZZZ"))
	fmt.Println(ZipString(""))
	fmt.Println(ZipString("ziiiiipMeee"))
	fmt.Println(ZipString("hhellloTthereYouuunggFelllas"))

}
