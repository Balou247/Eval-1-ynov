package main

import "github.com/01-edu/z01"

func countchar (s string, c rune) {
	var cpt int = 0
	for i:= 0; i < len(s); i++ {
		if c == s(i) {
			cpt += 1
		}
	}
	return cpt
}
