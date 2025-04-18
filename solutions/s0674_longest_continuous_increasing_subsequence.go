/*
	https://leetcode.com/problems/longest-continuous-increasing-subsequence/description

	Given an unsorted array of integers nums, return the length of the longest continuous increasing subsequence
	(i.e. subarray). The subsequence must be strictly increasing.

	A continuous increasing subsequence is defined by two indices l and r (l < r) such that it is [nums[l],
	nums[l + 1], ..., nums[r - 1], nums[r]] and for each l <= i < r, nums[i] < nums[i + 1].
*/

package solutions

func findLengthOfLCIS(nums []int) int {
	maxLen := 0
	c := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			c++
		} else {
			maxLen = max(maxLen, c)
			c = 1
		}
	}

	return max(maxLen, c)
}
