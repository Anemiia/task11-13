package main

import "unicode"

func CountVowels(str string) int {
	vowels := map[rune]bool{
		'а': true, 'е': true, 'ё': true, 'и': true, 'о': true,
		'у': true, 'ы': true, 'э': true, 'ю': true, 'я': true,
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}

	count := 0

	for _, char := range str {
		if vowels[unicode.ToLower(char)] {
			count++
		}
	}

	return count
}