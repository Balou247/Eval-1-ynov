package main

import (
	"github.com/01-edu/z01"
	"fmt"
	)
	
func RepeatAlpha(s string) string {
	var phrase str = ""
	for i := 0; i < ASCIINumber(s) - 97; i++ {
		z01.PrintRune(s)
	}
	return str
}

func main() {
	fmt.Println(piscine.RepeatAlpha("abc"))
	fmt.Println(piscine.RepeatAlpha("Choumi."))
	fmt.Println(piscine.RepeatAlpha(""))
	fmt.Println(piscine.RepeatAlpha("abacadaba 01!"))
}
