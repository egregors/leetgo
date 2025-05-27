/*
	https://leetcode.com/problems/largest-divisible-subset/description/

	Given a set of distinct positive integers nums, return the largest subset
		answer such that every pair (answer[i], answer[j]) of elements in this subset
		satisfies:

		answer[i] % answer[j] == 0, or
		answer[j] % answer[i] == 0

	If there are multiple solutions, return any of them.
*/

package solutions

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	maxLen, maxIdx := 1, 0
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = []int{nums[i]}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if len(dp[j])+1 > len(dp[i]) {
					dp[i] = append([]int(nil), dp[j]...)
					dp[i] = append(dp[i], nums[i])
				}
				if l := len(dp[i]); l > maxLen {
					maxLen = l
					maxIdx = i
				}
			}
		}
	}
	return dp[maxIdx]
}
