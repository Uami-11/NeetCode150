package main

import "strings"

func spellchecker(wordlist []string, queries []string) []string {
	var result []string

	exactMatch := make(map[string]bool)
	wordMatch := make(map[string]string)
	vowelMatch := make(map[string]string)

	for _, word := range wordlist {
		exactMatch[word] = true
		if _, exists := wordMatch[strings.ToLower(word)]; !exists {
			wordMatch[strings.ToLower(word)] = word
		}

		if _, exists := vowelMatch[replaceVowels(word)]; !exists {
			vowelMatch[replaceVowels(word)] = word
		}
	}

	for _, query := range queries {
		if exactMatch[query] {
			result = append(result, query)
		} else if word, exists := wordMatch[strings.ToLower(query)]; exists {
			result = append(result, word)
		} else if word, exists := vowelMatch[replaceVowels(query)]; exists {
			result = append(result, word)
		} else {
			result = append(result, "")
		}
	}

	return result
}

func replaceVowels(word string) string {
	word = strings.ToLower(word)
	var result strings.Builder
	for _, char := range word {
		if strings.ContainsRune("aeiou", char) {
			result.WriteRune('*')
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}
