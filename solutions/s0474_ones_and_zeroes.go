/*
	https://leetcode.com/problems/ones-and-zeroes/

	You are given an array of binary strings strs and two integers m and n.

	Return the size of the largest subset of strs such that there are at most m 0's
	and n 1's in the subset.

	A set x is a subset of a set y if all elements of x are also elements of y.
*/

package solutions

func findMaxForm(strs []string, m, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		count := countZeroesOnes(s)
		for zeroes := m; zeroes >= count[0]; zeroes-- {
			for ones := n; ones >= count[1]; ones-- {
				dp[zeroes][ones] = Maximum(1+dp[zeroes-count[0]][ones-count[1]], dp[zeroes][ones])
			}
		}
	}
	return dp[m][n]
}

func countZeroesOnes(s string) []int {
	c := make([]int, 2)
	for _, ch := range s {
		c[ch-'0']++
	}
	return c
}
