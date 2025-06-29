/*
	https://leetcode.com/problems/valid-word-abbreviation

	A string can be abbreviated by replacing any number of non-adjacent, non-empty
		substrings with their lengths. The lengths should not have leading zeros.

	For example, a string such as "substitution" could be abbreviated as (but not
		limited to):

		"s10n" ("s ubstitutio n")
		"sub4u4" ("sub stit u tion")
		"12" ("substitution")
		"su3i1u2on" ("su bst i t u ti on")
		"substitution" (no substrings replaced)

	The following are not valid abbreviations:

		"s55n" ("s ubsti tutio n", the replaced substrings are adjacent)
		"s010n" (has leading zeros)
		"s0ubstitution" (replaces an empty substring)

	Given a string word and an abbreviation abbr, return whether the string matches
		the given abbreviation.

	A substring is a contiguous non-empty sequence of characters within a string.
*/

package solutions

import (
	"strconv"
	"strings"
	"unicode"
)

func validWordAbbreviation(word string, abbr string) bool {
	if len(abbr) > len(word) {
		return false
	}

	var wi, ai int

	word = strings.ReplaceAll(word, " ", "")

	for wi < len(word) && ai < len(abbr) {
		a, b := word[wi], abbr[ai]

		if !unicode.IsDigit(rune(b)) {
			if a != b {
				return false
			}
			wi++
			ai++
		} else {
			if b == '0' {
				return false
			}
			cnt := []rune{}
			for unicode.IsDigit(rune(b)) {
				cnt = append(cnt, rune(b))
				ai++
				if ai > len(abbr)-1 {
					break
				}
				b = abbr[ai]
			}
			n, err := strconv.Atoi(string(cnt))
			if err != nil {
				panic(err)
			}
			wi += n
		}
	}

	return wi == len(word) && ai == len(abbr)
}
