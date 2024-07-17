package main

import (
	"fmt"
	//"strconv"
	//"sort"
	//"os"
	// "sort"
	//"strconv"
)
func main(){
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk (slice []int, size int){
	// if size > len(slice){
	// 	size = len(slice)
	// }
	if size == 0{
		fmt.Println()
		return
	}
	var result [][]int
	for i := 0; i < len(slice); i = i+size{
		end := i + size
		if end > len(slice){
			end = len(slice)
		}
		chunk := slice[i:end]
		result = append(result, chunk)
	}
fmt.Println(result)
}