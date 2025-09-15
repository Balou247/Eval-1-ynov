package main

import "github.com/01-edu/z01"

func RepeatAlpha(s string) string {
	var phrase str = ""
	for i := 0; i < ASCIINumber(s) - 97; i++ {
		z01.PrintRune(s)
	}
	return str
}
