package main

import (
	"unicode"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && (!unicode.IsDigit(rune(s[left])) && !unicode.IsLetter(rune(s[left]))) {
			left++
		}
		for left < right && (!unicode.IsDigit(rune(s[right])) && !unicode.IsLetter(rune(s[right]))) {
			right--
		}
		if unicode.ToUpper(rune(s[left])) != unicode.ToUpper(rune(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}
