/*
	https://leetcode.com/problems/verifying-an-alien-dictionary/

	In an alien language, surprisingly, they also use English lowercase letters,
		but possibly in a different order.
	The order of the alphabet is some permutation of lowercase letters.

	Given a sequence of words written in the alien language, and the order of the
		alphabet, return true if and only
	if the given words are sorted lexicographically in this alien language.
*/

package solutions

import "sort"

func isAlienSorted(words []string, order string) bool {
	m := make(map[rune]rune)
	for i, b := range order {
		m[b] = rune('a' + byte(i))
	}
	for i, ch := range words {
		runes := []rune(ch)
		for i, r := range runes {
			runes[i] = m[r]
		}
		words[i] = string(runes)
	}
	return sort.StringsAreSorted(words)
}
