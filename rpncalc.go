package main
import(
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){
	if len(os.Args) != 2{
		fmt.Println("ERROR")
		return
	}
	rpn(os.Args[1])
}

func rpn(str string){
	elems := strings.Split(str, " ")
	elems = deleteSpaces(elems)
	values := []int{}

	for _, val := range elems{
		num, err := strconv.Atoi(val)
		if err == nil {
			values = append(values, num)
			continue
		}

		n := len(values)
		if n < 2{
			fmt.Println("Error")
			return
		}
		switch val{
		case "+":
			values[n-2] += values[n-1]
			values = values[:n-1]
		case "-":
			values[n-2] -= values[n-1]
			values = values[:n-1]
		case "*":
			values[n-2] *= values[n-1]
			values = values[:n-1]
		case "/":
			values[n-2] /= values[n-1]
			values = values[:n-1]
		case "%":
			values[n-2] += values[n-1]
			values = values[:n-1]
		}
	}
	if len(values) == 1{
		fmt.Println(values[0])
		return
	}
	fmt.Println("Error")
	return
}

func deleteSpaces(s []string) []string{
	var res []string

	for _, char := range s{
		if char != ""{
			res = append(res, char)
		}
	}
	return res
}