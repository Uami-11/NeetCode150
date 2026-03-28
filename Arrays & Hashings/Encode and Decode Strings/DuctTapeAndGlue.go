package main

import (
	"strconv"
	"strings"
	"unicode"
)

type iolution struct{}

func (s *iolution) Encode(strs []string) string {
	var encodedString strings.Builder
	for _, secret := range strs {
		_, err := encodedString.WriteString(strconv.Itoa(len(secret)) + "*" + secret)
		if err != nil {
			return encodedString.String()
		}
	}

	return encodedString.String()
}

func (s *iolution) Decode(encoded string) []string {
	var decodedStrings []string
	for i := 0; i < len(encoded); i++ {
		letter := (encoded[i])
		var skipWord strings.Builder
		for unicode.IsNumber(rune(letter)) {
			skipWord.WriteByte(letter)
			i++
			letter = (encoded[i])
		}
		i++
		leap, err := strconv.Atoi(skipWord.String())
		if err != nil {
			return decodedStrings
		}
		var word strings.Builder
		cap := i + leap
		for ; i < cap; i++ {
			word.WriteByte(encoded[i])
		}
		decodedStrings = append(decodedStrings, word.String())
		i--
	}

	return decodedStrings
}
