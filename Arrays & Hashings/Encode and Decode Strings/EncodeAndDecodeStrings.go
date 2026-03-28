package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encodedString strings.Builder
	for _, str := range strs {
		fmt.Fprintf(&encodedString, "%d*%s", len(str), str)
	}

	return encodedString.String()
}

func (s *Solution) Decode(encoded string) []string {
	var decodedStrings []string
	for i := 0; i < len(encoded); {
		sep := strings.IndexByte(encoded[i:], '*') + i
		length, _ := strconv.Atoi(encoded[i:sep])
		start := sep + 1
		decodedStrings = append(decodedStrings, encoded[start:start+length])
		i = start + length
	}

	return decodedStrings
}
