/*
	https://leetcode.com/problems/palindromic-substrings/

	Given a string s, return the number of palindromic substrings in it.

	A string is a palindrome when it reads the same backward as forward.

	A substring is a contiguous sequence of characters within the string.
*/

package solutions

func countSubstrings(s string) int {
	if s == "" {
		return 0
	}
	count := 0

	for i := range s {
		extendPal(s, i, i, &count)
		extendPal(s, i, i+1, &count)
	}

	return count
}

func extendPal(s string, lo, hi int, count *int) {
	for lo >= 0 && hi < len(s) && s[lo] == s[hi] {
		*count++
		lo--
		hi++
	}
}
