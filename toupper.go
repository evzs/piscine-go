package piscine

func ToUpper(s string) string {

	var upper string

	for _, up := range s {

		if up >= 97 && up <= 122 {

			upper += string(rune(up - 32))

		} else {
			upper += string(up)
		}
	}
	return upper
}
