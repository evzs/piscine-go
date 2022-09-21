package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(IsPrime(7))
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(-5))
}

func IsPrime(nb int) bool {

	if nb <= 1 {
		return false
	}

	for x := 2; x < nb; x++ {
		if nb%x == 0 {
			return false
		}
	}
	return true
}
