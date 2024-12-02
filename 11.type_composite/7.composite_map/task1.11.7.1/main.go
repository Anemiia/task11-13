package main

import (
	"fmt"
	"strings"
)

func countWordOccurrences(text string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		counts[word]++
	}

	return counts
}

func main() {
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	occurrences := countWordOccurrences(text)

	for word, count := range occurrences {
		fmt.Printf("%s: %d\n", word, count)
	}
}