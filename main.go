package main

import (
	"fmt"

	piscine "piscine/3"
)

func main() {
	fmt.Println(piscine.DigitLen(100, 10))
	fmt.Println(piscine.DigitLen(100, 2))
	fmt.Println(piscine.DigitLen(-100, 16))
	fmt.Println(piscine.DigitLen(100, -1))
}

// NotDecimal("51.3+95.9") == "513+959\n" instead of "51.3+95.9\n"
