package main

import "strings"

func createUniqueText(text string) string {
	counts := make(map[string]int)
	words := strings.Fields(text)

	var uniqueWords []string
	for _, word := range words {
		counts[word]++
		if counts[word] == 1 {
			uniqueWords = append(uniqueWords, word)
		}
	}

	return strings.Join(uniqueWords, " ")
}