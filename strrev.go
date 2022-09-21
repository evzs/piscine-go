package piscine

import (
	"fmt"
)

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) (result string) {
	for _, r := range s {
		result = string(r) + result
	}
	return
}
