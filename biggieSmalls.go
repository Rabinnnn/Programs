package main

import (
	"fmt"
)

func main() {
	// Define smallest and largest values for integers
	var (
		smalls int = -(1 << 63) // Smallest possible value for int
		biggie int = 1<<63 - 1  // Largest possible value for int
	)

	fmt.Println("smalls:", smalls)
	fmt.Println("biggie:", biggie)
}

//alternative
// * import (
  //  "fmt"
 //   "math"
//)
//
//func main() {
  //  smalls := math.MinInt64  // Smallest possible integer value (int64)
  //  biggie := math.MaxInt64  // Largest possible integer value (int64)
  //  
   // fmt.Println("smalls:", smalls)
 //   fmt.Println("biggie:", biggie)
//} 
