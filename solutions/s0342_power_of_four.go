/*
	https://leetcode.com/problems/power-of-four/

	Given an integer n, return true if it is a power of four. Otherwise, return
		false.

	An integer n is a power of four, if there exists an integer x such that n ==
		4x.
*/

package solutions

func isPowerOfFour(n int) bool {
	if n < 0 || n&(n-1) != 0 {
		return false
	}
	return n%3 == 1
}
