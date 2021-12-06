/*
	https://leetcode.com/problems/longest-palindrome/

	Given a string s which consists of lowercase or uppercase letters, return the length
	of the longest palindrome that can be built with those letters.

	Letters are case sensitive, for example, "Aa" is not considered a palindrome here.
*/

package solutions

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, ch := range s {
		m[ch]++
	}

	var res string
	for k, v := range m {
		for v > 1 {
			res = string(k) + res + string(k)
			v -= 2
			m[k] -= 2
		}
	}

	if len(res)%2 == 0 {
		for k, v := range m {
			if v > 0 {
				res = res[len(res)/2:] + string(k) + res[:len(res)/2]
				break
			}
		}
	}

	return len(res)
}
