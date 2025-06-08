package piscine

func FishAndChips(n int) string {
	res := ""
	if n < 0 {
		return "error: number is negative"
	}

	if n%2 == 0 && n%3 == 0 {
		res = "fish and chips"
	} else if n%3 == 0 {
		res = "chips"
	} else if n%2 == 0 {
		res = "fish"
	} else {
		res = "error: non divisible"
	}
	return res
}
