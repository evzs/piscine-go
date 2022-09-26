package piscine

func NRune(s string, n int) rune {
	arr := []rune(s)
	counter := 0

	for range arr {

		counter++
	}

	if n > 0 && n <= counter {

		return arr[n-1]
	}
	return 0
}
