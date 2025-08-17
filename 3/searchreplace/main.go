package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	str := args[0]
	letter := []rune(args[1])[0]
	replace := []rune(args[2])[0]

	res := ""

	for _, v := range str {
		if v == letter {
			v = replace
		}
		res += string(v)
	}
	res += "\n"

	for _, v := range res {
		z01.PrintRune(v)
	}

}
