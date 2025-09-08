package main

import "github.com/01-edu/z01"

func CountAlpha(s string) int { 
	var cpt int = 0 
	for i := 0; i < len(s); i++ { 
		cpt += 1 
	}
	return cpt 
}
