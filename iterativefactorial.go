package piscine

import (
	"fmt"
)

func main() {
	arg := 9
	fmt.Println(IterativeFactorial(arg))
}

func IterativeFactorial(nb int) int {

	result := 1
	for i := 1; i <= nb; i++ {
		result = result * i
	}
	return result

}
