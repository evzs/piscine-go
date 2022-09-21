package piscine

import "github.com/01-edu/z01"

func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}

func PrintWordsTables(a []string) {
	for _, v := range a {
		for _, a := range v {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
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
