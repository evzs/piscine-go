package piscine

import (
	"fmt"
)

func main() {
	s := "Hello 78 World! 4555 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}

func AlphaCount(s string) int {
	counter := 0
	for _, nb := range s {
		if (nb >= 'a' && nb <= 'z') || (nb >= 'A' && nb <= 'Z') {
			counter++
		}
	}
	return counter
}
