/*
	https://leetcode.com/problems/find-words-containing-character/description/

		You are given a 0-indexed array of strings words and a character x.

	Return an array of indices representing the words that contain the character x.

	Note that the returned array may be in any order.
*/

package solutions

import "strings"

func findWordsContaining(words []string, x byte) []int {
	res := make([]int, 0, len(words))

	for i, s := range words {
		if strings.ContainsRune(s, rune(x)) {
			res = append(res, i)
		}
	}

	return res
}
