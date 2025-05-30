/*
	https://leetcode.com/problems/is-subsequence/

	Given two strings s and t, return true if s is a subsequence of t, or false
		otherwise.

	A subsequence of a string is a new string that is formed from the original
		string by
	deleting some (can be none) of the characters without disturbing the relative
		positions
	of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while
		"aec" is not).
*/

package solutions

func isSubsequence(s, t string) bool {
	slow, fast := 0, 0
	for fast < len(t) && slow < len(s) {
		if s[slow] == t[fast] {
			slow++
		}
		fast++
	}
	return slow == len(s)
}
