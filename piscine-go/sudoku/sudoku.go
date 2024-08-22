package main

import (
	"fmt"
	"os"
)

const size = 9

func main() {
	if len(os.Args) != size+1 {
		fmt.Println("Error")
		return
	}

	var sudoku [size][size]int
	for i, arg := range os.Args[1:] {
		if len(arg) != size {
			fmt.Println("Error")
			return
		}
		for j, char := range arg {
			if char == '.' {
				sudoku[i][j] = 0
			} else if char >= '1' && char <= '9' {
				sudoku[i][j] = int(char - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	if solveSudoku(&sudoku) {
		printSudoku(&sudoku)
	} else {
		fmt.Println("Error")
	}
}

func solveSudoku(sudoku *[size][size]int) bool {
	row, col := findUnassignedLocation(sudoku)
	if row == -1 && col == -1 {
		return true // No unassigned location; puzzle solved
	}

	for num := 1; num <= 9; num++ {
		if isSafe(sudoku, row, col, num) {
			sudoku[row][col] = num

			if solveSudoku(sudoku) {
				return true // Successfully solved the puzzle
			}

			sudoku[row][col] = 0 // Backtrack
		}
	}

	return false
}

func findUnassignedLocation(sudoku *[size][size]int) (int, int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if sudoku[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1 // No unassigned location found
}

func isSafe(sudoku *[size][size]int, row, col, num int) bool {
	return !usedInRow(sudoku, row, num) &&
		!usedInCol(sudoku, col, num) &&
		!usedInBox(sudoku, row-row%3, col-col%3, num)
}

func usedInRow(sudoku *[size][size]int, row, num int) bool {
	for i := 0; i < size; i++ {
		if sudoku[row][i] == num {
			return true
		}
	}
	return false
}

func usedInCol(sudoku *[size][size]int, col, num int) bool {
	for i := 0; i < size; i++ {
		if sudoku[i][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(sudoku *[size][size]int, startRow, startCol, num int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku[startRow+i][startCol+j] == num {
				return true
			}
		}
	}
	return false
}

func printSudoku(sudoku *[size][size]int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(sudoku[i][j])
			if j < size-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
