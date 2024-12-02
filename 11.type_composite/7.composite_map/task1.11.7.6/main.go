package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Word struct {
	Word string
	Pos  int
}

// filterWords Фильтрует текст, заменяя цензурные и повторяющиеся слова
func filterWords(text string, censorMap map[string]string) string {
	sentences := splitSentences(text)

	if len(sentences) > 1 {
		for i, sentence := range sentences {
			sentences[i] = filterWords(sentence, censorMap)
		}
		return strings.Join(sentences, " ")
	}

	words := strings.Fields(text)

	uniqueWords := make(map[string]Word)

	for i, word := range words {
		oldWord := strings.Trim(word, "!,")
		lowerOld := strings.ToLower(oldWord)

		//Если слово содержится в карте цензурных слов, то
		//Замена слова на значение из карты, используя CheckUpper
		if newWord, ok := censorMap[lowerOld]; ok {
			words[i] = CheckUpper(oldWord, newWord)
			lowerOld = strings.ToLower(strings.Trim(words[i], "!,"))
		}

		//Если слово не содержится в карте уникальных слов, то (для проверки ключа в карте уникальных слов, используйте функцию strings.ToLower)
		//Добавление слова в карту уникальных слов
		if _, exists := uniqueWords[lowerOld]; !exists {
			uniqueWords[lowerOld] = Word{Word: words[i], Pos: i}
			continue
		}
		// Если слово содержится в карте уникальных слов, то нужно его очистить
		words[i] = ""
	}

	return WordsToSentence(words)
}

// WordsToSentence Удаляет пустые слова из слайса и объединяет их в предложение, добавляя в конце восклицательный знак
func WordsToSentence(words []string) string {
	filtered := make([]string, 0, len(words))

	for _, word := range words {
		if word != "" {
			filtered = append(filtered, word)
		}
	}

	return strings.ReplaceAll(strings.Join(filtered, " ")+"!", "!!", "!")
}

// CheckUpper Проверяет, нужно ли заменять первую букву на заглавную
func CheckUpper(old, new string) string {
	if len(old) == 0 || len(new) == 0 {
		return new
	}

	chars := []rune(old)

	if unicode.IsUpper(chars[0]) {
		runes := []rune(new)
		new = string(append([]rune{unicode.ToUpper(runes[0])}, runes[1:]...))
	}

	return new
}

// splitSentences Разделяет текст на предложения
func splitSentences(message string) []string {
	// Создание регулярного выражения для поиска знаков препинания
	originSentences := strings.Split(message, "!")
	var orphan string
	var sentences []string

	for i, sentence := range originSentences {
		words := strings.Split(sentence, " ")

		if len(words) == 1 {
			if len(orphan) > 0 {
				orphan += " "
			}

			orphan += words[0] + "!"
			continue
		}

		if orphan != "" {
			originSentences[i] = strings.Join([]string{orphan, " ", sentence}, " ") + "!"
			orphan = ""
		}

		sentences = append(sentences, originSentences[i])
	}

	return sentences
}

func main() {
	text := "Внимание! Внимание! Покупай срочно срочно крипту только у нас! Биткоин лайткоин эфир по низким ценам! Беги, беги, успевай стать финансово независимым с помощью крипты! Крипта будущее финансового мира!"
	censorMap := map[string]string{
		"крипта":   "фрукты",
		"крипту":   "фрукты",
		"крипты":   "фруктов",
		"биткоин":  "яблоки",
		"лайткоин": "яблоки",
		"эфир":     "яблоки",
	}

	filteredText := filterWords(text, censorMap)
	fmt.Println(filteredText) // Внимание! Покупай срочно фрукты только у нас! Яблоки по низким ценам! Беги, успевай стать финансово независимым с помощью фруктов! Фрукты будущее финансового мира!
}