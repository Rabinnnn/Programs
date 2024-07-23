package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))

	fmt.Println(HashCode("14 Avril 1912"))

}

func HashCode(dec string) string {
	res := ""
	val := 0
	for _, char := range dec {
		val = (int(char) + len(dec)) % 127
		if val < 32 || val > 126 {
			val = val + 33
		}
		res += string(val) //string(val) returns a string of one rune NOT digit

	}

	return res
}
