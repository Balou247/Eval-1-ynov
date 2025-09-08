package main

import "github.com/01-edu/z01"

func PrintReverseAlphabet () {
	var num int = 122
	for i := 0; i < 26; i++ {
		z01.PrintRune(rune(num + i))
	}
	z01.PrintRune('/n')
}
