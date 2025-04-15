/*
	https://leetcode.com/problems/valid-palindrome/description

	A phrase is a palindrome if, after converting all uppercase letters into lowercase
	letters and removing all non-alphanumeric characters, it reads the same forward and backward.
	Alphanumeric characters include letters and numbers.

	Given a string s, return true if it is a palindrome, or false otherwise.
*/

package solutions

func isPalindrome125(s string) bool {
	isValid := func(ch byte) bool {
		return (ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9')
	}
	toLowel := func(ch byte) byte {
		if 'A' <= ch && ch <= 'Z' {
			ch = ch - 'A' + 'a'
		}
		return ch
	}

	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isValid(s[l]) {
			l++
		}
		for l < r && !isValid(s[r]) {
			r--
		}
		if toLowel(s[l]) != toLowel(s[r]) {
			return false
		}
		l++
		r--
	}

	return true
}
