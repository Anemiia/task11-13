package main

func countBytes(s string) int {
	return len(s)
}

func countSymbols(s string) int {
	return len([]rune(s))
}

