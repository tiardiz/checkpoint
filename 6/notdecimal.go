package piscine

func isChar(ch rune) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}
func isValidSymbol(ch rune) bool {
	return (ch >= '0' && ch <= '9') || ch == '-' || ch == '.'
}

func NotDecimal(dec string) string {
	if len(dec) == 0 {
		return "\n"
	}
	res := ""
	for i, v := range dec {
		if isChar(v) {
			return dec + "\n"
		} else if !isValidSymbol(v) {
			return dec + "\n"
		} else if i == 0 && v == '-' {
			res += string(v)
		} else if i == 0 && v == '0' && dec[i+1] == '.' {
			continue
		} else if v != '.' {
			res += string(v)
		}
	}
	return trimLeadingZeros(res) + "\n"
}
func trimLeadingZeros(s string) string {
	i := 0
	for i < len(s) && s[i] == '0' {
		i++
	}
	if i == len(s) {
		return "0"
	}
	return s[i:]
}
