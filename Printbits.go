package main
import(
	"fmt"
	"os"
	"strconv"
)
func main() {
fmt.Println(Printbits(os.Args[1]))

}

func Printbits(s string)string{
	num, _ := strconv.Atoi(s)
	binary := ""
	
	for i := 7; i >= 0; i--{
		if num & (1<<i) != 0{
			binary += "1"
		}else{
			binary += "0"
		}
	}
	
	return binary
}