/*
	https://leetcode.com/problems/power-of-two/

	Given an integer n, return true if it is a power of two. Otherwise, return false.

	An integer n is a power of two, if there exists an integer x such that n == 2x
*/

package solutions

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return (n & (-n)) == n
}
