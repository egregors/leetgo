/*
	https://leetcode.com/problems/palindrome-partitioning/

	Given a string s, partition s such that every substring of the partition is a
		palindrome.
	Return all possible palindrome partitioning of s.
*/

package solutions

func partition131(s string) [][]string {
	isPalindrome := func(s string, lo, hi int) bool {
		for lo < hi {
			if s[lo] != s[hi] {
				return false
			}
			lo++
			hi--
		}
		return true
	}

	var res [][]string
	var bt func(start int, curr []string, s string)
	bt = func(start int, curr []string, s string) {
		if start >= len(s) {
			cp := make([]string, len(curr))
			copy(cp, curr)
			res = append(res, cp)
		}
		for end := start; end < len(s); end++ {
			if isPalindrome(s, start, end) {
				curr = append(curr, s[start:end+1])
				bt(end+1, curr, s)
				curr = curr[:len(curr)-1]
			}
		}
	}
	bt(0, []string{}, s)

	return res
}
