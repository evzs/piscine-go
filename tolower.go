package piscine

import (
	"fmt"
)

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}

func ToLower(s string) string {

	var lower string

	for _, low := range s {

		if low >= 65 && low <= 90 {

			lower += string(rune(low + 32))

		} else {
			lower += string(low)
		}
	}
	return lower
}
