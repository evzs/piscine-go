package piscine

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
