/*
	https://leetcode.com/problems/palindrome-permutation/description/

	Given a string s, return true if a permutation of the string could form a and false otherwise.
*/

package solutions

func canPermutePalindrome(s string) bool {
	m := make([]int, 26)
	for _, ch := range s {
		m[ch-'a']++
	}

	var odd int
	if len(s)%2 != 0 {
		odd = 1
	}

	for _, c := range m {
		if c != 0 && c%2 != 0 {
			odd--
			if odd < 0 {
				return false
			}
		}
	}

	return true
}
