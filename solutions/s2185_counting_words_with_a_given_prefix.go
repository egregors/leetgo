/*
	https://leetcode.com/problems/counting-words-with-a-given-prefix

	You are given an array of strings words and a string pref.

	Return the number of strings in words that contain pref as a prefix.

	A prefix of a string s is any leading contiguous substring of s.
*/

package solutions

import "strings"

func prefixCount(words []string, pref string) int {
	cnt := 0
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			cnt++
		}
	}
	return cnt
}
