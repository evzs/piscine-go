package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(Index("Hello! ", "l"))
	fmt.Println(Index("Salut! ", "alu"))
	fmt.Println(Index("0la! ", "h01"))
}

func Index(s string, toFind string) int {
	str := []rune(s)
	x := []rune(toFind)
	y := 0
	z := 0
	for range str {
		y++
	}
	for range x {
		z++
	}
	for i := 0; i <= y-z; i++ {
		if toFind == s[i:i+z] {
			return (i)
		}
	}
	return -1
}
