package main

import (
	//"fmt"
	"unicode"
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

// func PrintMemory(a [10]byte) {
// 	str := ""
// 	for i, nbr := range a {
// 		fmt.Printf("%.2x", nbr)

// 		if ((i+1)%4 == 0 && i != 0) || i == len(a)-1 {
// 			fmt.Println()
// 		} else {
// 			fmt.Print(" ")
// 		}

// 		if nbr >= 33 && nbr <= 126 {
// 			str += string(rune(nbr))
// 		} else {
// 			str += "."
// 		}
// 	}
// 	fmt.Println(str + strings.Repeat(".", 10-len(a)))
// }


func PrintMemory(arr [10]byte) {
	output := ""
	// Step 1: Display bytes in hexadecimal format
	for i, b := range arr {
		// Print the hex value of each byte
		//fmt.Printf("%02x ", b)
		output += intToHex(int(b))
		// Add a new line after every 4 bytes for readability
		if (i+1)%4 == 0  || i == len(arr)-1{
			output += "\n"
		}else{
			output += " "
		}
	}

	// Step 2: Display ASCII graphic characters or replace non-printable ones with dots
	for _, b := range arr {
		// Check if the byte is a printable ASCII graphic character
		if unicode.IsGraphic(rune(b)) && b >= 32 && b <= 126 {
			//fmt.Printf("%c", b)
			output += string(b)
		} else {
			// Replace non-printable characters with dots
			output += "."
		}
	}
	// End with a new line
//	fmt.Println()
for _, char := range output{
	z01.PrintRune(char)
}
z01.PrintRune('\n')
}

func intToHex(n int) string {
	if n == 0 {
		return "00"
	}

	hexChars := "0123456789abcdef"
	result := ""
	for n > 0 {
		result = string(hexChars[n%16]) + result
		n /= 16
	}
	return result
}