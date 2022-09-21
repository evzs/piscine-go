package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(IsUpper("hello"))
	fmt.Println(IsUpper("hello!"))
	fmt.Println(IsUpper("HELLO"))
}

func IsUpper(s string) bool {
	for _, low := range s {
		if low <= 'a' || low >= 'z' {
			return false
		}
	}
	return true
}
