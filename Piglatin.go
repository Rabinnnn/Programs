package main
import(
	"fmt"
	"os"
//	"strconv"
)
func main() {
fmt.Println(Pigy(os.Args[1]))

}

func Pigy(s string)string{
	result := ""
	for i := 0; i < len(s); i++{
		if s[0] == 'a' || s[0] == 'e' || s[0] == 'i' || s[0] == 'o' || s[0] == 'u' ||
		s[0] == 'A' || s[0] == 'E' || s[0] == 'I' || s[0] == 'O' || s[0] == 'U'{
			result = s + "ay"
		}else{
			for i, char := range s{
				if !(char == 'a' || char == 'e' ||char == 'i' ||char == 'o' ||char == 'u' ||
				char == 'A' || char == 'E' ||char == 'I' ||char == 'O' ||char == 'U') {
					result = s[i:] + s[0:i] + "ay"
				}
			}
		}
	}
	count := 0
	for _, char := range s{
		if char == 'a' || char == 'e' ||char == 'i' ||char == 'o' ||char == 'u' ||
				char == 'A' || char == 'E' ||char == 'I' ||char == 'O' ||char == 'U'{
					count = count + 1
		}
	}
	if count == 0{
		result = "No vowels"
	}

	return result
}