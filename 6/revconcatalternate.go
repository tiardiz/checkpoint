package piscine

func RevConcatAlternate(num1, num2 []int) []int {
	res := []int{}
	if len(num1) == len(num2) {
		for i := 0; i < len(num1); i++ {
			res = append(res, num1[(len(num1)-1)-i])
			res = append(res, num2[(len(num2)-1)-i])
		}
	}
	if len(num1) == 0 {
		for i := 0; i < len(num2); i++ {
			// res = append(res, num1[(len(num2)-1)-i])
			res = append(res, num2[(len(num2)-1)-i])
		}
	} else if len(num2) == 0 {
		for i := 0; i < len(num1); i++ {
			// res = append(res, num1[(len(num2)-1)-i])
			res = append(res, num1[(len(num1)-1)-i])
		}
	}
	if len(num1) > len(num2) {
		rev := num1[(len(num1)-len(num2))-1:]
		for i := 0; i < len(rev); i++ {
			res = append(res, rev[(len(rev)-1)-i])
		}
		for i := 0; i < len(num2); i++ {
			//Подумай еще!!!!!
			res = append(res, rev[(len(rev)-1)-i])
			res = append(res, num2[(len(num2)-1)-i])
		}

	}
	return res
}
