package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, nb := range s {
		if (nb >= 'a' && nb <= 'z') || (nb >= 'A' && nb <= 'Z') {
			counter++
		}
	}
	return counter
}
