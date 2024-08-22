package piscine

// import "fmt"

func Abort(a, b, c, d, e int) int {
	// Put the five integers into a slice for easier sorting
	numbers := []int{a, b, c, d, e}

	// Sort the numbers in ascending order
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				// Swap if the current element is greater than the next
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	// Calculate the median
	median := numbers[2] // The middle element after sorting

	return median
}

/*
func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
} */
