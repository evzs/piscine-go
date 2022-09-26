package piscine

func IsUpper(s string) bool {
	for _, up := range s {
		if up <= 'A' || up >= 'Z' {
			return false
		}
	}
	return true
}
