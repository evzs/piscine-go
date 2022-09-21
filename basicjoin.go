package piscine

import (
	"fmt"
)

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}

func BasicJoin(elems []string) string {
	nb := ""
	for _, r := range elems {
		nb += r
	}
	return nb
}
