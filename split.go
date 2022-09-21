package piscine

import (
	"fmt"
)

func main() {
	s := "WeshHAMonHAGars"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	ln := 0

	for i := range sep {
		ln = i + 1
	}

	ln2 := 0
	for i := range s {
		ln2 = i + 1
	}

	for i := 0; i < ln2-ln; i++ {
		if s[i:i+ln] == sep {
			s = s[:i] + " " + s[i+ln:]
			ln2 -= ln
		}
	}

	return SplitWhiteSpaces(s)
}

func SplitWhiteSpaces(s string) []string {
	var sresult []string

	a := ""

	for i := 0; i < len(s); i++ {

		a = ""

		for s[i] != ' ' && s[i] != '\n' && s[i] != 9 {

			a += string(s[i])

			if i == len(s)-1 {

				break

			}

			i++

		}

		if a != "" {

			sresult = append(sresult, a)

		}

	}

	return sresult
}
