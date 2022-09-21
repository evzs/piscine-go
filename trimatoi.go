package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45 "))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World! "))
	fmt.Println(TrimAtoi("sd+x1fa2W3 s4"))
	fmt.Println(TrimAtoi("sd -x1fa 2W3 s4 "))
	fmt.Println(TrimAtoi("sdx1 -fa2W3 s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}

func TrimAtoi(s string) int {
	x := true
	y := false
	z := 0
	var nb rune = 48
	for _, r := range s {
		if r == '-' && x {
			y = true
		}
		if r >= nb && r <= 57 {
			z *= 10
			z += int(r) - 48
			x = false
		}
	}
	if y == true {
		return -z
	}
	return z
}
