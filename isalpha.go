package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(IsAlpha("thisisatest"))
	fmt.Println(IsAlpha("this is a test"))
	fmt.Println(IsAlpha("777!"))
	fmt.Println(IsAlpha("7 7 7"))
}

func IsAlpha(s string) bool {
	for _, alpha := range s {
		if alpha < 48 {
			return false
		} else if alpha > 57 && alpha < 65 {
			return false
		} else if alpha > 90 && alpha < 97 {
			return false
		} else if alpha > 122 {
			return false
		}
	}
	return true
}
