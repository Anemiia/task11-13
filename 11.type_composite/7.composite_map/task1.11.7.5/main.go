package main

import (
	"fmt"
	"strings"
)

func filterSentence(sentence string, filter map[string]bool) string {
	counts := make(map[string]int)
	words := strings.Fields(sentence)

	var uniqueWords []string

	for _, word := range words {
		counts[word]++
		if !filter[word] { // проверяется что word в мапе filter является false
			uniqueWords = append(uniqueWords, word)
		}
	}

	return strings.Join(uniqueWords, " ")
}

func main() {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}

	filteredSentence := filterSentence(sentence, filter)
	fmt.Println(filteredSentence)
}