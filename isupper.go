package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
	fmt.Println(IsUpper("hello"))
}

func IsUpper(s string) bool {
	for _, up := range s {
		if up <= 'A' || up >= 'Z' {
			return false
		}
	}
	return true
}
