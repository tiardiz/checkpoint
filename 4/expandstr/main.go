package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		return
	}

	input := os.Args[1]
	res := ""
	for i := range input {

		if isAplha(input[i]) {
			res += string(input[i])
		} else if i > 0 && input[i] == ' ' && isAplha(input[i-1]) {
			res += "   "
		} else if i > len(os.Args[1]) {
			break
		}
	}

	for _, v := range delSpacesEnd(res) {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
func delSpacesEnd(str string) string {
	end := len(str) - 1
	for end >= 0 && str[end] == ' ' {
		end--
	}
	return str[:end+1]
}

func isAplha(ch byte) bool {
	if ch > 32 && ch < 127 {
		return true
	}
	return false
}
