package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}

func MakeRange(min, max int) []int {

	if min >= max {

		return nil
	}
	size := max - min
	answer := make([]int, size)

	for i := 0; i < size; i++ {
		answer[i] = min
		min++
	}

	return answer
}
