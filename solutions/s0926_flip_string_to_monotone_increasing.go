/*
	https://leetcode.com/problems/flip-string-to-monotone-increasing/description/

	A binary string is monotone increasing if it consists of some number of 0's
	(possibly none), followed by some number of 1's (also possibly none).

	You are given a binary string s. You can flip s[i] changing it from 0 to 1
	or from 1 to 0.

	Return the minimum number of flips to make s monotone increasing.
*/

package solutions

func minFlipsMonoIncr(s string) int {
	res, n := 0, 0
	for _, ch := range s {
		if ch == '0' {
			res = Minimum(n, res+1)
		} else {
			n++
		}
	}
	return res
}
