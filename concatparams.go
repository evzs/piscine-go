package piscine

import (
	"fmt"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}

func ConcatParams(args []string) string {
	sresult := ""

	for i := 0; i < len(args); i++ {

		sresult += args[i] + "\n"
	}
	return sresult
}
