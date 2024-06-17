package main

import(
	"fmt"
)

func main(){
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	Bubble(result)
	fmt.Println(result)

}

func Bubble(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}