/*
	https://leetcode.com/problems/house-robber/

	You are a professional robber planning to rob houses along a street.
	Each house has a certain amount of money stashed, the only constraint stopping you
	from robbing each of them is that adjacent houses have security systems connected and
	it will automatically contact the police if two adjacent houses were broken into on the
	same night.

	Given an integer array nums representing the amount of money of each house,
	return the maximum amount of money you can rob tonight without alerting the police.
*/

package solutions

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[n], dp[n-1] = 0, nums[n-1]

	for i := n - 2; i >= 0; i-- {
		dp[i] = max(dp[i+1], dp[i+2]+nums[i])
	}

	return dp[0]
}
