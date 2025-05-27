/*
	https://leetcode.com/problems/minimum-average-difference/

	You are given a 0-indexed integer array nums of length n.

	The average difference of the index i is the absolute difference between the
		average of
	the first i + 1 elements of nums and the average of the last n - i - 1
		elements.
	Both averages should be rounded down to the nearest integer.

	Return the index with the minimum average difference. If there are multiple
		such indices,
	return the smallest one.

	Note:

		The absolute difference of two numbers is the absolute value of their
			difference.
		The average of n elements is the sum of the n elements divided (integer
			division) by n.
		The average of 0 elements is considered to be 0.

*/

package solutions

import "math"

func minimumAverageDifference(nums []int) int {
	n := len(nums)
	ans := -1
	minAvgDiff := math.MaxInt

	prefixSum := make([]int, n+1)
	suffixSum := make([]int, n+1)

	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	for i := n - 1; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + nums[i]
	}

	for i := 0; i < n; i++ {
		lAvg := prefixSum[i+1]
		lAvg /= i + 1

		rAvg := suffixSum[i+1]
		if i != n-1 {
			rAvg /= n - i - 1
		}

		currDiff := Abs(lAvg - rAvg)
		if currDiff < minAvgDiff {
			minAvgDiff = currDiff
			ans = i
		}
	}
	return ans
}
