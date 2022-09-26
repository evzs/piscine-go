package piscine

func IsNumeric(s string) bool {
	for _, numeric := range s {
		if numeric < 48 || numeric > 57 {
			return false
		}
	}
	return true
}
