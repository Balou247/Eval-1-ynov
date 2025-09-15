package main

func RepeatAlpha(s string) string {
	var phrase str = ""
	for i := 0; i < ASCIINumber(s) - 97; i++ {
		phrase += "s"
	}
	return phrase
}
