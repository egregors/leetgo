/*
	https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/

	Given two string arrays word1 and word2, return true if the two arrays represent the same string,
	and false otherwise.

	A string is represented by an array if the array elements concatenated in order forms the string.
*/

package solutions

import "strings"

func arrayStringsAreEqual(word1, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}
