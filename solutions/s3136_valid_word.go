/*
	https://leetcode.com/problems/valid-word/

		A word is considered valid if:

		It contains a minimum of 3 characters.
		It contains only digits (0-9), and English letters (uppercase and lowercase).
		It includes at least one vowel.
		It includes at least one consonant.

	You are given a string word.

	Return true if word is valid, otherwise, return false.

	Notes:

		'a', 'e', 'i', 'o', 'u', and their uppercases are vowels.
		A consonant is an English letter that is not a vowel.
*/

package solutions

import (
	"strings"
	"unicode"
)

func isValid3136(word string) bool {
	if len(word) < 3 {
		return false
	}
	var l, c bool

	for _, ch := range strings.ToLower(word) {
		if !unicode.IsDigit(ch) && !unicode.IsLetter(ch) {
			return false
		}

		if unicode.IsLetter(ch) {
			if vowels[ch] {
				l = true
			} else {
				c = true
			}
		}
	}

	return l && c
}

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}
