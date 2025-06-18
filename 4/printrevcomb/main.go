package main

import "github.com/01-edu/z01"

func main() {
	for i := '9'; i >= '0'; i-- {
		for j := '8'; j >= '0'; j-- {
			for k := '7'; k >= '0'; k-- {
				if i > j && j > k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
				} else {
					continue
				}
				if i != '2' || j != '1' || k != '0' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
