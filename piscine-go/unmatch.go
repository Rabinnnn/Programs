package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	for 0 < len(a) {
		if len(a)-1 != 0 {
			if a[0] == a[1] {
				a = a[2:]
			} else {
				return a[0]
			}
		} else {
			return a[0]
		}
	}
	return -1
}
