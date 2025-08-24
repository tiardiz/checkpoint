package piscine

func ThirdTimeIsACharm(str string) string {
	if len(str) < 3 {
		return "\n"
	}
	res := ""
	runeStr := []rune(str)
	for i := 0; i < len(runeStr); i++ {
		if (i+1)%3 == 0 {
			res += string(runeStr[i])
		}
	}
	return res + "\n"
}

// 1 2 3 4 5 6 7 8 9
//     2     5     8
