package main

import (
	"os"

	"github.com/01-edu/z01"
)

func toLower(ch rune) rune {
	if ch < 97 {
		return ch + 32
	}
	return ch
}

func isVowel(ch rune) bool {
	if toLower(ch) == 'a' || toLower(ch) == 'o' || toLower(ch) == 'i' || toLower(ch) == 'e' || toLower(ch) == 'u' || toLower(ch) == 'y' {
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 || len(args) >= 2 {
		return
	}
	str := args[0]
	start := ""
	end := ""
	idx := 0
	for i := range str {
		if isVowel(toLower(rune(str[i]))) {
			idx = i
			break
		}
	}
	start = str[idx:]
	end = str[:idx] + "ay\n"
	res := start + end

	if res != "" {
		for _, ch := range res {
			z01.PrintRune(ch)
		}
	} else if res == "" && str != "" {
		for _, ch := range "No vowels\n" {
			z01.PrintRune(ch)
		}
	}
}
