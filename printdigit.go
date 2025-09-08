package main

func PrintReverseAlphabet () {
	var num int = 48
	for i := 0; i < 10; i++ {
		z01.PrintRune(rune(num + i))
	}
	z01.PrintRune('/n')
}
