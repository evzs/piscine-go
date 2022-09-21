package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(IsNumeric("thisisatest"))
	fmt.Println(IsNumeric("this is a test"))
	fmt.Println(IsNumeric("777!"))
	fmt.Println(IsNumeric("7 7 7"))
	fmt.Println(IsNumeric("7"))
}

func IsNumeric(s string) bool {
	for _, numeric := range s {
		if numeric < 48 || numeric > 57 {
			return false
		}
	}
	return true
}
