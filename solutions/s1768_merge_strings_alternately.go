/*
	https://leetcode.com/problems/merge-strings-alternately/

	You are given two strings word1 and word2. Merge the strings by
	adding letters in alternating order, starting with word1. If a
	string is longer than the other, append the additional letters
	onto the end of the merged string.

	Return the merged string.
*/

package solutions

func mergeAlternately(word1, word2 string) string {
	if word1 == "" {
		return word2
	}
	if word2 == "" {
		return word1
	}
	return string(word1[0]) + string(word2[0]) + mergeAlternately(word1[1:], word2[1:])
}
