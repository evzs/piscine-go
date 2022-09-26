package piscine

func TrimAtoi(s string) int {
	x := true
	y := false
	z := 0
	var nb rune = 48
	for _, r := range s {
		if r == '-' && x {
			y = true
		}
		if r >= nb && r <= 57 {
			z *= 10
			z += int(r) - 48
			x = false
		}
	}
	if y == true {
		return -z
	}
	return z
}
