package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, value := range a {
		result = append(result, f(value))
	}
	return result
}

/*
func Map(f func(int) bool, a []int) []bool {
	b := make([]bool, len(a))
	for i, v := range a {
		b[i] = f(v)
	}
	return b
}
*/
