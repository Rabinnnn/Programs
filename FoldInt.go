package main

import (
	"fmt"
	//"strconv"
	//"sort"
	//"os"
	// "sort"
	//"strconv"
)


func FoldInt(f func(int, int) int, a []int, ac int){
	res := ac
	for i := 0; i < len(a); i++{
		res = f(res,a[i])
	}
	fmt.Println(res)
}

func main() {
	Mul := func(acc int, cur int) int {
		return acc * cur
	}
	Add := func(acc int, cur int) int {
		return acc + cur
	}
	Sub := func(acc int, cur int) int {
		return acc - cur
	}
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}