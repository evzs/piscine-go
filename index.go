package piscine

func Index(s string, toFind string) int {
	str := []rune(s)
	x := []rune(toFind)
	y := 0
	z := 0
	for range str {
		y++
	}
	for range x {
		z++
	}
	for i := 0; i <= y-z; i++ {
		if toFind == s[i:i+z] {
			return (i)
		}
	}
	return -1
}
