package piscine

func ReverseMenuIndex(menu []string) []string {
	// Get the length of the menu slice
	length := len(menu)

	// Create a new slice with the same length
	reversedMenu := make([]string, length)

	// Copy elements from the original slice to the new slice in reverse order
	for i := 0; i < length; i++ {
		reversedMenu[i] = menu[length-1-i]
	}

	return reversedMenu
}
