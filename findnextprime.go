package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
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

func FindNextPrime(nb int) int {
	if nb < 3 {
		return 2
	}
	for x := nb; ; x++ {
		IsPrime := true
		for y := 2; y*y <= x; y++ {
			if x%y == 0 {
				IsPrime = false
				break
			}
		}
		if IsPrime {
			return x
		}
	}
}
