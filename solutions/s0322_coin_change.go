/*
	https://leetcode.com/problems/coin-change/

	You are given an integer array coins representing coins of different denominations
	and an integer amount representing a total amount of money.

	Return the fewest number of coins that you need to make up that amount. If that
	amount of money cannot be made up by any combination of the coins, return -1.

	You may assume that you have an infinite number of each kind of coin.
*/

package solutions

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = Minimum(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
