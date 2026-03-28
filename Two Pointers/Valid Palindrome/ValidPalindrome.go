package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	str := strings.ToUpper(s)
	pointer := 1
	for i := 0; i < len(str)/2; {
		end := len(str) - pointer
		if str[i] != str[end] {
			if !unicode.IsDigit(rune(str[i])) && !unicode.IsLetter(rune(str[i])) {
				i++
			} else if !unicode.IsDigit(rune(end)) && !unicode.IsLetter(rune(str[end])) {
				pointer++
			} else {
				return false
			}
		} else {
			i++
			pointer++
		}
	}
	return true
}
