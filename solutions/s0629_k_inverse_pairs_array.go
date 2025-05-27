/*
	https://leetcode.com/problems/k-inverse-pairs-array/

	For an integer array nums, an inverse pair is a pair of integers [i, j] where
	0 <= i < j < nums.length and nums[i] > nums[j].

	Given two integers n and k, return the number of different arrays consist of
	numbers from 1 to n such that there are exactly k inverse pairs. Since the
		answer
	can be huge, return it modulo 109 + 7.
*/

package solutions

//nolint:revive // wtf?
func kInversePairs(n, k int) int {
	dp := make([]int, k+1)
	M := 1000000007
	for i := 1; i <= n; i++ {
		temp := make([]int, k+1)
		temp[0] = 1
		for j := 1; j <= k; j++ {
			var tmp int
			if j-i >= 0 {
				tmp = dp[j-i]
			}

			val := (dp[j] + M - tmp) % M
			temp[j] = (temp[j-1] + val) % M
		}
		dp = temp
	}
	var tmp int
	if k > 0 {
		tmp = dp[k-1]
	}
	return (dp[k] + M - tmp) % M
}
