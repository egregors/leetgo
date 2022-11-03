/*
	https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/

	You are given an array of strings words. Each element of words consists of two lowercase English letters.

	Create the longest possible palindrome by selecting some elements from words and concatenating them in any
	order. Each element can be selected at most once.

	Return the length of the longest palindrome that you can create. If it is impossible to create any palindrome,
	return 0.

	A palindrome is a string that reads the same forward and backward.
*/

package solutions

// longestPalindrome2131 should call longestPalindrome2131 to pass Leetcode tests
func longestPalindrome2131(words []string) int {
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}
	var (
		res int
		mid bool
	)

	swap := func(s string) string {
		bs := []byte(s)
		bs[0], bs[1] = bs[1], bs[0]
		return string(bs)
	}

	for w, c := range m {
		w2 := swap(w)
		if w != w2 {
			if c2, ok := m[w2]; ok && c > 0 {
				res += 4 * Minimum(c, c2)
				delete(m, w)
				delete(m, w2)
			}
		} else {
			if rem := c % 2; rem != 0 && !mid {
				res += rem * 2
				mid = true
			}
			res += c / 2 * 4
			delete(m, w)
		}
	}

	return res
}
