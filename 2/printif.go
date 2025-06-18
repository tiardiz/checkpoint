package piscine

func PrintIf(str string) string {
	if len(str) >= 3 {
		str = "G\n"
	} else if str == "" {
		str = "G\n"
	} else {
		str = "Invalid Input\n"
	}
	return str
}
