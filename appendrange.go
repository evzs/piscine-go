package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(AppendRange(7, 16))
	fmt.Println(AppendRange(8, 1))
}

func AppendRange(min, max int) []int {

	var s []int
	for i := min; i < max; i++ {
		s = append(s, i)
	}

	if min >= max {

		return nil
	}
	return s
}
