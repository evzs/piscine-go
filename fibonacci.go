package piscine

import (
	"fmt"
)

func main() {
	arg1 := 7
	fmt.Println(Fibonacci(arg1))
}

func Fibonacci(index int) int {
	if index < 2 {
		return index
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
