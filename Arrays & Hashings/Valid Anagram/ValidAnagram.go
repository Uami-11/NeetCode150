func isAnagram(s string, t string) bool {
	freq := make(map[rune]int)

	for _, char := range s {
		freq[char]++
	}
	for _, char := range t {
		freq[char]--
	}

	for _, value := range freq {
		if value != 0 {
			return false
		}
	}

	return true
}
