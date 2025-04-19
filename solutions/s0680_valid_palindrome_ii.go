/*
	https://leetcode.com/problems/valid-palindrome-ii/description

	Given a string s, return true if the s can be palindrome after deleting at most one character from it.
*/

package solutions

func validPalindrome(s string) bool {
	check := func(s string, l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	for i := 0; i < len(s)>>1; i++ {
		if s[i] != s[len(s)-1-i] {
			return check(s, i, len(s)-2-i) || check(s, i+1, len(s)-1-i)
		}
	}

	return true
}
