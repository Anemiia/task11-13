package main

func ReplaceSymbols(str string, old, new rune) string {
	runes := []rune(str)
	for i, char := range runes {
		if char == old {
			runes[i] = new
		}
	}
	return string(runes)
}