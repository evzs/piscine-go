package piscine

func ConcatParams(args []string) string {
	sresult := ""

	for i := 0; i < len(args); i++ {

		sresult += args[i] + "\n"
	}
	return sresult
}
