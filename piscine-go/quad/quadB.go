package main

import "fmt"

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 1; i <= y; i++ {
		if i == 1 || i == y {
			for j := 1; j <= x; j++ {
				if j == 1 && i == 1 {
					fmt.Print("/")
				} else if j == x && i == 1 {
					fmt.Print("\\")
				} else if j == 1 && i == y {
					fmt.Print("\\")
				} else if j == x && i == y {
					fmt.Print("/")
				} else {
					fmt.Print("*")
				}
			}
		} else {
			for j := 1; j <= x; j++ {
				if j == 1 || j == x {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
