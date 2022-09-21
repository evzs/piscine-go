package piscine

import (
	"fmt"
)

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}

func Join(strs []string, sep string) string {
	a := 0
	for range strs {
		a++
	}

	b := ""
	for i := 0; i < a; i++ {
		b += strs[i]
		if i != a-1 {
			b += sep
		}
	}
	return b
}
