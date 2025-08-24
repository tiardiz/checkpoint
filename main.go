package main

import (
	"fmt"
	piscine "piscine/6"
)

func main() {
	fmt.Print(piscine.NotDecimal("0.1"))
	fmt.Print(piscine.NotDecimal("174.2"))
	fmt.Print(piscine.NotDecimal("0.1255"))
	fmt.Print(piscine.NotDecimal("1.20525856"))
	fmt.Print(piscine.NotDecimal("-0.0f00d00"))
	fmt.Print(piscine.NotDecimal(""))
	fmt.Print(piscine.NotDecimal("-19.525856"))
	fmt.Print(piscine.NotDecimal("1952"))
	fmt.Print(piscine.NotDecimal("00.02"))
	fmt.Print(piscine.NotDecimal("51.3+95.9"))
	fmt.Println(rune('.'))
}

// NotDecimal("51.3+95.9") == "513+959\n" instead of "51.3+95.9\n"
