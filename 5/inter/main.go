package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	charMap := make(map[rune]bool)
	for _, char := range str2 {
		charMap[char] = true
	}

	seen := make(map[rune]bool)

	for _, char := range str1 {
		if charMap[char] && !seen[char] {
			z01.PrintRune(char)
			seen[char] = true
		}
	}
	z01.PrintRune('\n')
}
