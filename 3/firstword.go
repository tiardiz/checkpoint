package piscine

func FirstWord(s string) string {
	res := ""
	if len(s) == 0 {
		return "\n"
	}
	for _, v := range s {
		if v != ' ' {
			res += string(v)
		} else if len(res) == 0 && v == ' ' {
			continue
		} else {
			res += string('\n')
			break
		}
	}
	return res
}
