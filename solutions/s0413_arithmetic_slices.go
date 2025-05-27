/*
	https://leetcode.com/problems/arithmetic-slices/

	An integer array is called arithmetic if it consists of at least three elements
		and if
	the difference between any two consecutive elements is the same.

		For example, [1,3,5,7,9], [7,7,7,7], and [3,-1,-5,-9] are arithmetic
			sequences.

	Given an integer array nums, return the number of arithmetic subarrays of nums.

	A subarray is a contiguous subsequence of the array.
*/

package solutions

func numberOfArithmeticSlices(nums []int) int {
	dp := make([]int, len(nums))
	sum := 0

	for i := 2; i < len(dp); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = 1 + dp[i-1]
			sum += dp[i]
		}
	}

	return sum
}
