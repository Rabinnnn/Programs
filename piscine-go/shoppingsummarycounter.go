package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	currentItem := ""

	for _, char := range str {
		if char == ' ' {
			summary[currentItem]++
			currentItem = ""
		} else {
			currentItem += string(char)
		}
	}

	summary[currentItem]++
	return summary
}
