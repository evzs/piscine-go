package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(Sqrt(80))
	fmt.Println(Sqrt(40))
}

func Sqrt(nb int) int {
	for i := 1; nb/i >= i; i++ {
		if nb/i == i {
			return i
		}
	}
	return 0
}
