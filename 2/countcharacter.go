package piscine

func CountChar(str string, c rune) int {
	cnt := 0
	for _, v := range str {
		if v == c {
			cnt++
		}
	}
	return cnt
}
