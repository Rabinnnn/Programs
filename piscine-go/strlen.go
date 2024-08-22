package piscine

func StrLen(s string) int {
	runeCount := 0

	for range s {
		runeCount++
	}

	return runeCount
}
