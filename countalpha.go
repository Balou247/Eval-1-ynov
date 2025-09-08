package main

import "unicode"

func CountAlpha(s string) int {
	cpt := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			cpt++
		}
	}
	return cpt
}