package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) < 2 {
		return
	}

	input := os.Args[1]
	input = low(input)

	for _, v := range input {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
func low(input string) string {
	res := ""
	for i := range input {

		if isLowerCase(input[i]) {
			res += string(input[i])
		} else if isUpperCase(input[i]) {
			res += string(input[i] + 32)
			// fmt.Println(res)
		} else if isLowerCase(input[i]) && input[i+1] == ' ' {
			res += string(input[i] - 32)
		} else if i > len(os.Args[1]) {
			break
		} else {
			res += string(input[i])
		}
	}
	return res
}

func isLowerCase(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

func isUpperCase(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}

// func isAlpha(ch byte) bool {
// 	return isLowerCase(ch) || isUpperCase(ch)
// }
