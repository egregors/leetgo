/*
	https://leetcode.com/problems/valid-anagram/

	Given two strings s and t, return true if t is an anagram of s, and false otherwise.
*/

package solutions

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	xs := make(map[rune]int)
	for _, ch := range s {
		xs[ch]++
	}

	for _, ch := range t {
		xs[ch]--
	}

	for _, v := range xs {
		if v != 0 {
			return false
		}
	}
	return true
}
