package piscine

func IsSorted(f func(int, int) bool, a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if !f(a[i], a[i+1]) {
			return false
		}
	}
	return true
}

func f(a, b int) bool {
	return a <= b
}
