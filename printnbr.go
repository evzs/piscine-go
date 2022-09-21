package piscine

import "github.com/01-edu/z01"

func main() {
	PrintNbr(-123)
	PrintNbr(123)
	z01.PrintRune('\n')
	PrintNbr(383838)
}

func PrintNum(num int) {
	x := '0'
	if num == 0 {
		z01.PrintRune('0')
		return
	}
	for y := 1; y <= num%10; y++ {
		x++
	}
	for y := -1; y >= num%10; y-- {
		x++
	}
	if num/10 != 0 {
		PrintNum(num / 10)
	}
	z01.PrintRune(x)
}

func PrintNbr(n int) {

	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNum(n)
}
