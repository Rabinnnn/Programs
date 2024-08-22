package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	arg1 := os.Args[1]
	operator := os.Args[2]
	arg2 := os.Args[3]

	result := evaluate(arg1, operator, arg2)
	if result != "" {
		writeToStdout(result)
	}  
}

func evaluate(arg1, operator, arg2 string) string {
	switch operator {
	case "+":
		return add(arg1, arg2)
	case "-":
		return subtract(arg1, arg2)
	case "*":
		return multiply(arg1, arg2)
	case "/":
		return divide(arg1, arg2)
	case "%":
		return modulo(arg1, arg2)
	default:
		return ""
	}
}

func add(arg1, arg2 string) string {
	// Implement addition logic here
	return ""
}

func subtract(arg1, arg2 string) string {
	// Implement subtraction logic here
	return ""
}

func multiply(arg1, arg2 string) string {
	// Implement multiplication logic here
	return ""
}

func divide(arg1, arg2 string) string {
	// Implement division logic here
	return ""
}

func modulo(arg1, arg2 string) string {
	// Implement modulo logic here
	return ""
}

func writeToStdout(result string) {
	os.Stdout.WriteString(result)
}
