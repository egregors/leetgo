/*
	https://leetcode.com/problems/count-number-of-homogenous-substrings/description/

	Given a string s, return the number of homogenous substrings of s. Since the answer may be too large,
	return it modulo 109 + 7.

	A string is homogenous if all the characters of the string are the same.

	A substring is a contiguous sequence of characters within a string.
*/

package solutions

func countHomogenous(s string) int {
	var res, c int

	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] == s[i-1] {
			c++
		} else {
			c = 1
		}

		res = (res + c) % 1_000_000_007
	}

	return res
}
