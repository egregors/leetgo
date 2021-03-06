/*
	https://leetcode.com/problems/remove-palindromic-subsequences/

	You are given a string s consisting only of letters 'a' and 'b'. In a
	single step you can remove one palindromic subsequence from s.

	Return the minimum number of steps to make the given string empty.

	A string is a subsequence of a given string if it is generated by deleting
	some characters of a given string without changing its order. Note that a
	subsequence does not necessarily need to be contiguous.

	A string is called palindrome if is one that reads the same backward as well
	as forward.
*/

package solutions

func isPalindrome1332(s string) bool {
	lo, hi := 0, len(s)-1
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}

func removePalindromeSub(s string) int {
	if s == "" {
		return 0
	}
	if isPalindrome1332(s) {
		return 1
	}
	return 2
}
