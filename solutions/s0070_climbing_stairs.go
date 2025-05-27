/*
	https://leetcode.com/problems/climbing-stairs/

	You are climbing a staircase. It takes n steps to reach the top.

	Each time you can either climb 1 or 2 steps. In how many distinct ways can you
		climb to the top?
*/

package solutions

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
