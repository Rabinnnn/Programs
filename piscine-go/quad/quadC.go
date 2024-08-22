package main

import "fmt"

func QuadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 {
				if j == 1 {
					fmt.Print("A")
				} else if j == x {
					fmt.Print("A")
				} else {
					fmt.Print("B")
				}
			} else if i == y {
				if j == 1 {
					fmt.Print("C")
				} else if j == x {
					fmt.Print("C")
				} else {
					fmt.Print("B")
				}
			} else {
				if j == 1 || j == x {
					fmt.Print("B")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
func main() {
	// Use default dimensions for demonstration
	QuadC(1, 1)
}
