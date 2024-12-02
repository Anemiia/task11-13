package main

import (
	"fmt"
	"strings"
)

func CountWordsInText(txt string, words []string) map[string]int {
	counts := make(map[string]int)
	text := strings.Fields(txt)

	for _, word := range text {
		for _, slovo := range words {
			if slovo == strings.ToLower(word) {
				counts[slovo]++
			}
		}
	}

	return counts
}

func main() {
	txt := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. 
        Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. 
        Praesent et diam eget libero egestas mattis sit amet vitae augue.`
	words := []string{"sit", "amet", "lorem"}

	result := CountWordsInText(txt, words)

	fmt.Println(result) // map[amet:2 lorem:1 sit:3]
}
