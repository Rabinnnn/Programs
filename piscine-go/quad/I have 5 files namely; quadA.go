I have 5 files namely; quadA.go, quadB.go, quadC.go, quadD.go, and quadE.go. They are in the same directory. As an example, here is the content of quadA.go

package piscine

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 || i == y-1 {
				if j == 0 || j == x-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			} else {
				if j == 0 || j == x-1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}


I want you to create a program in golang called quadchecker. It should take a string as an argument and display the name of the matching quad and its dimensions.
If the argument is not a quad the program should print Not a quad function.

All answers must end with a newline ('\n').

If there is more than one quad matches, the program must display them all alphabetically ordered and separated by a ||.

Remember, there are 5 quad files which means there are also 5 quad functions.
Here is the expected usage and output to guide you. Make sure you adhere to the instructions and that when your code is executed as shown below it 
will give the exact output.
