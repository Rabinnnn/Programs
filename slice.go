package main

import "fmt"

func Slice(a []string, nbrs ...int) []string {
	// Handle the case where no integers are provided
	if len(nbrs) == 0 {
		return []string{}
	}

	// Handle the case where only one integer is provided
	start := nbrs[0]
	end := len(a) // Default end is the length of the slice

	// Handle negative indices
	if start < 0 {
		start += len(a)
	}
	if len(nbrs) > 1 {
		end = nbrs[1]
		if end < 0 {
			end += len(a)
		}
	}

	// Ensure start and end are within bounds
	if start < 0 {
		start = 0
	}
	if end > len(a) {
		end = len(a)
	}

	// Return the slice from start to end
	if start > end {
		return []string(nil) // If start is greater than end, return an empty slice
	}
	return a[start:end]
}

func main(){
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n", Slice(a, 1))
    fmt.Printf("%#v\n", Slice(a, 2, 4))
    fmt.Printf("%#v\n", Slice(a, -3))
    fmt.Printf("%#v\n", Slice(a, -2, -1))
    fmt.Printf("%#v\n", Slice(a, 2, 0))
}
