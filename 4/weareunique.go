package piscine

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	cnt := 0
	str1Map := map[rune]int{}
	str2Map := map[rune]int{}

	for _, ch1 := range str1 {
		str1Map[ch1]++
	}
	for _, ch2 := range str2 {
		str2Map[ch2]++
	}
	for key := range str1Map {
		if _, exist := str2Map[key]; !exist {
			cnt++
		}
	}
	for key := range str2Map {
		if _, exist := str1Map[key]; !exist {
			cnt++
		}
	}
	return cnt
}
