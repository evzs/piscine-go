package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(RecursivePower(7, 3))
}

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}
