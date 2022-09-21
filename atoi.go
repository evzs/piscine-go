package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
	ar := []rune(s)
	n := len(s)
	ans := 0
	sign := 1
	for i := 0; i < n; i++ {
		if ar[i] == '-' {
			sign = -1
		} else if ar[i] == '+' {
			sign = 1
		} else if ar[i] < '0' || ar[i] > '9' {
			return 0
		} else {
			ans *= 10
			ans += int(ar[i]) - '0'
		}

	}
	return ans * sign
}
