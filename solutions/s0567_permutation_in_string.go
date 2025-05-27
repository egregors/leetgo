/*
	https://leetcode.com/problems/permutation-in-string/

	Given two strings s1 and s2, return true if s2 contains a permutation of s1, or
		false otherwise.

	In other words, return true if one of s1's permutations is the substring of s2.
*/

package solutions

func isPermut(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1m, s2m := make(map[uint8]int, 26), make(map[uint8]int, 26)
	for i := 0; i < len(s1); i++ {
		s1m[s1[i]]++
		s2m[s2[i]]++
	}

	for k, v := range s1m {
		if s2m[k] != v {
			return false
		}
	}

	return true
}

func checkInclusion(s1, s2 string) bool {
	left, right := 0, len(s1)
	for right <= len(s2) {
		if isPermut(s1, s2[left:right]) {
			return true
		}
		left++
		right++
	}
	return false
}
