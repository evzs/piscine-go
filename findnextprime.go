package piscine

func FindNextPrime(nb int) int {
	if nb < 3 {
		return 2
	}
	for x := nb; ; x++ {
		IsPrime := true
		for y := 2; y*y <= x; y++ {
			if x%y == 0 {
				IsPrime = false
				break
			}
		}
		if IsPrime {
			return x
		}
	}
}
