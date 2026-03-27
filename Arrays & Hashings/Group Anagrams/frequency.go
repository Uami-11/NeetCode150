package main

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]int][]string)

	for _, word := range strs {
		var freq [26]int
		for _, char := range word {
			freq[int(char)-int('a')]++
		}
		anagrams[freq] = append(anagrams[freq], word)
	}

	var grouped [][]string
	for _, anagramSlice := range anagrams {
		grouped = append(grouped, anagramSlice)
	}

	return grouped
}
