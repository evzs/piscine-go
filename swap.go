package piscine

import (
	"fmt"
)

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}
