package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 2 {
		return
	}
	str := args[0]
	runeStr := []rune(str)
	word := ""
	words := []string{}
	for i := 0; i < len(runeStr); i++ {
		if runeStr[i] != ' ' {
			word += string(runeStr[i])
		} else if runeStr[i] == ' ' {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	revWords := []string{}
	for i := len(words) - 1; i >= 0; i-- {
		revWords = append(revWords, words[i])
	}
	for i, word := range revWords {
		for _, ch := range word {
			z01.PrintRune(ch)
		}
		if i != len(revWords)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
