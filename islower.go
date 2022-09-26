package piscine

func IsLower(s string) bool {
	for _, low := range s {
		if low <= 'a' || low >= 'z' {
			return false
		}
	}
	return true
}
