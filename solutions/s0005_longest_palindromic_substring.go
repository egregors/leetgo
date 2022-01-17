/*
	https://leetcode.com/problems/longest-palindromic-substring/

	Given a string s, return the longest palindromic substring in s.
*/

package solutions

func isPalindrome(s string) bool {
	mid := len(s) >> 1
	for i := 0; i < mid; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome5(s string) string {
	if isPalindrome(s) {
		return s
	}

	res := ""
	for i := 0; i < len(s)-1; i++ {
		lo, hi := i, i+1
		for {
			// a -> a[a]a
			if lo > 0 && hi < len(s) && isPalindrome(s[lo-1:hi+1]) {
				lo--
				hi++
				continue
			}

			// a ->  [a]a
			if hi < len(s) && isPalindrome(s[lo:hi+1]) {
				hi++
				continue
			}

			// a -> a[a]
			if lo > 0 && isPalindrome(s[lo-1:hi]) {
				lo--
				continue
			}

			break
		}

		if hi-lo > len(res) {
			res = s[lo:hi]
		}
	}

	return res
}
