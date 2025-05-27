/*
	https://leetcode.com/problems/word-pattern/

	Given a pattern and a string s, find if s follows the same pattern.

	Here follow means a full match, such that there is a bijection between a letter
		in pattern and a non-empty word in s.
*/

package solutions

import (
	"strings"
)

func wordPattern(pattern, s string) bool {
	words := strings.Split(s, " ")
	wordsLen := len(words)

	if len(pattern) != wordsLen {
		return false
	}

	seenPattern := make(map[string]int, wordsLen)
	seenWords := make(map[string]int, wordsLen)

	for i := 0; i < wordsLen; i++ {
		ch := string(pattern[i])
		w := words[i]

		if _, ok := seenPattern[ch]; !ok {
			seenPattern[ch] = i
		}
		if _, ok := seenWords[w]; !ok {
			seenWords[w] = i
		}

		if seenPattern[ch] != seenWords[w] {
			return false
		}
	}

	return true
}
