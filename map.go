package piscine

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}

func Map(f func(int) bool, a []int) []bool {
	b := make([]bool, len(a))
	for i, v := range a {
		b[i] = f(v)
	}
	return b
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
