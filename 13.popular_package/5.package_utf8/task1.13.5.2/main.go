package main

import (
	"fmt"
	"unicode"
)

func countRussianLetters(s string) map[rune]int {
	counts := make(map[rune]int)
	runes := []rune(s)
	isRussianLetter := func(s rune) bool {
		return s >= 'а' && s <= 'я'
	}
	for _, simbol := range runes {
		simbol = unicode.ToLower(simbol)
		if isRussianLetter(simbol){
			counts[simbol]++
		}
	}
	return counts
}

func main() {
	result := countRussianLetters("Привет, мир!")
	for key, value := range result {
		fmt.Printf("%c: %d ", key, value) // в: 1 е: 1 т: 1 м: 1 п: 1 р: 2 и: 2
	}
}
