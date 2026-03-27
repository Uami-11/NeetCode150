package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, word := range strs {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedWords := string(runes)
		anagrams[sortedWords] = append(anagrams[sortedWords], word)
	}

	var grouped [][]string
	for _, anagramSlice := range anagrams {
		grouped = append(grouped, anagramSlice)
	}

	return grouped
}
