package piscine

func RepeatAlpha(s string) string {
	res := ""
	if s == "" {
		res = ""
	}

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			incr := int(ch) - 96
			for i := 1; i <= incr; i++ {
				res += string(ch)
			}
		} else if ch >= 'A' && ch <= 'Z' {
			inc := int(ch) - 64
			for i := 1; i <= inc; i++ {
				res += string(ch)
			}
		} else {

			res += string(ch)
		}
	}
	return res
}
