package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func isQuad(name string) bool {
	return strings.HasPrefix(name, "quad") && strings.HasSuffix(name, ".go")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: quadchecker <quad_name>")
		os.Exit(1)
	}

	quadName := os.Args[1]

	files := []string{"quadA.go", "quadB.go", "quadC.go", "quadD.go", "quadE.go"}

	for _, file := range files {
		if isQuad(file) {
			cmd := exec.Command("go", "run", file)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()

			if err != nil {
				fmt.Printf("[%s] Not a quad function\n", quadName)
				return
			}

			fmt.Printf("[%s] [%s]\n", quadName, file)
		}
	}
}
