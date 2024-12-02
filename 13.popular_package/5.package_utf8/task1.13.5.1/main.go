package main

func countUniqueUTF8Chars(s string) int {
	counts := make(map[rune]int)
	runes := []rune(s)

	for _, simbol := range runes {
		counts[simbol]++
	}
	return len(counts)
}
