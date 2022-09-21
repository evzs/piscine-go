package piscine

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune(NRune("additional test", 15))
	z01.PrintRune('\n')
}

func NRune(s string, n int) rune {
	arr := []rune(s)
	counter := 0

	for range arr {

		counter++
	}

	if n > 0 && n <= counter {

		return arr[n-1]
	}
	return 0
}
