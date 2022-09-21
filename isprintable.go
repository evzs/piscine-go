package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(IsPrintable("hello"))
	fmt.Println(IsPrintable("hello\n"))

}

func IsPrintable(s string) bool {

	tab := []rune(s)

	for _, printable := range tab {

		if printable >= 0 && printable < 32 {

			return false
		}
	}
	return true
}
