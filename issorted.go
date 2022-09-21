package piscine

import (
	"fmt"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

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
