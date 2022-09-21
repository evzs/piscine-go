package piscine

import "github.com/01-edu/z01"

func main() {

	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)

}

func IsNegative(x int) {

	if x <= 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
}
