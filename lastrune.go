package piscine

import (
	"github.com/01-edu/z01"
)

func main() {
	LastRune("testp")
}

func LastRune(s string) {
	var y rune

	for _, x := range s {
		y = x
	}
	z01.PrintRune(y)
}

/*	var y rune

	for i, x := range s {
		fmt.Println("Le premier élément est : ", string(x), "à l'index", i)
		y = x
	}
	fmt.Println("La dernière lettre est", y)*/
