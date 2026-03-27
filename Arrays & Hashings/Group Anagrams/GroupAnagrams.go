package main

import "slices"

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, word := range strs {
		runes := []rune(word)
		slices.Sort(runes)
		sortedWords := string(runes)
		anagrams[sortedWords] = append(anagrams[sortedWords], word)
	}

	var grouped [][]string
	for _, anagramSlice := range anagrams {
		grouped = append(grouped, anagramSlice)
	}

	return grouped
}
