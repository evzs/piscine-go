package piscine

import (
	"fmt"
)

func main() {
	arg := 7
	fmt.Println(RecursiveFactorial(arg))
}

func RecursiveFactorial(nb int) int {

	if nb == 1 {
		return 1
	}
	if nb > 1 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0

}
