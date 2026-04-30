package main

import (
	"strings"
	"unicode"
)

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	var seenWord bool
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsLetter(rune(s[i])) {
			seenWord = true
		}
		if s[i] == ' ' && seenWord {
			return len(s) - 1 - i
		}
	}

	return len(s)
}
