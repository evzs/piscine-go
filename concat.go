package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(Concat("Hello!", " How are you?"))
}

func Concat(a string, b string) string {
	concat := a + b
	return concat
}
