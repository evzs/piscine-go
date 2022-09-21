package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(IterativePower(7, 3))
}

func IterativePower(nb int, power int) int {

	if power < 0 {
		return 0
	} else if 1 == power {
		return nb
	} else if 0 == power {
		return 1
	} else {
		result := 1
		for i := 0; i < power; i++ {
			result *= nb
		}
		return result
	}
}