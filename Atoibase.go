package main

import(
	"fmt"
	//"os"
)

func main(){
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
    // Validate the base
    if len(base) < 2 {
        return 0
    }
    
    baseSet := make(map[byte]int)
    for i := 0; i < len(base); i++ {
        if base[i] == '+' || base[i] == '-' {
            return 0
        }
        if _, exists := baseSet[base[i]]; exists {
            return 0 // Duplicate characters in base
        }
        baseSet[base[i]] = i
    }
    
    // Convert string s based on base
    result := 0
    for i := 0; i < len(s); i++ {
        result = result * len(base) + baseSet[s[i]]
    }
    
    return result
}
/* ALTERNATIVE
func AtoiBase(s string, base string)int{
	result := 0

	if len(base) < 2{
		return 0
	}
	seen := make(map[rune]bool)
	baseM := make(map[rune]int)
	for i, char := range base{
		if char == '+' || char == '-'{
			return 0
		}
		if !seen[char]{
			seen[char] = true
		}else{
			return 0
		}
		baseM[char] = i
	}
	
	for _, char := range s{
		result = result * len(base) + baseM[char]
	}
	return result
} */
