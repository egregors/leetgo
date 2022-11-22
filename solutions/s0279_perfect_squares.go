/*
	https://leetcode.com/problems/perfect-squares/

	Given an integer n, return the least number of perfect square numbers that sum to n.

	A perfect square is an integer that is the square of an integer; in other words, it is the product of some
	integer with itself. For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.
*/

package solutions

import (
	"math"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = math.MaxInt
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = Minimum(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
