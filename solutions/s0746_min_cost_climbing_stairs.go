/*
	https://leetcode.com/problems/min-cost-climbing-stairs/

	You are given an integer array cost where cost[i] is the cost of ith step on a
		staircase.
	Once you pay the cost, you can either climb one or two steps.

	You can either start from the step with index 0, or the step with index 1.

	Return the minimum cost to reach the top of the floor.
*/

package solutions

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + Minimum(dp[i-1], dp[i-2])
	}

	if dp[len(dp)-1] < dp[len(dp)-2] {
		return dp[len(dp)-1]
	}

	return dp[len(dp)-2]
}
