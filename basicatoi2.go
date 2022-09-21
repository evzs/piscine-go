package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}

func BasicAtoi2(s string) int {
	ar := []rune(s)
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		if ar[i] < '0' || ar[i] > '9' {
			return 0
		} else {
			ans *= 10
			ans += int(ar[i]) - '0'
		}

	}
	return ans
}
