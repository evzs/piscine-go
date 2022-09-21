package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(Compare("Hello! ", "Hello! "))
	fmt.Println(Compare("Salut! ", "lut! "))
	fmt.Println(Compare("0la!", "01"))
}

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return +1
}
