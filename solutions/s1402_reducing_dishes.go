/*
	https://leetcode.com/problems/reducing-dishes/description/

	A chef has collected data on the satisfaction level of his n dishes. Chef can
		cook any dish in 1 unit of time.

	Like-time coefficient of a dish is defined as the time taken to cook that dish
		including previous dishes multiplied
	by its satisfaction level i.e. time[i] * satisfaction[i].

	Return the maximum sum of like-time coefficient that the chef can obtain after
		dishes preparation.

	Dishes can be prepared in any order and the chef can discard some dishes to get
		this maximum value.
*/

package solutions

import "sort"

func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	n := len(satisfaction)

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+2)
		for j := 0; j < n+2; j++ {
			dp[i][j] = 0
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= n; j++ {
			dp[i][j] = Maximum(satisfaction[i]*j+dp[i+1][j+1], dp[i+1][j])
		}
	}

	return dp[0][1]
}
