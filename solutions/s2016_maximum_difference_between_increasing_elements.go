/*
	https://leetcode.com/problems/maximum-difference-between-increasing-elements

	Given a 0-indexed integer array nums of size n, find the maximum difference
		between nums[i] and nums[j] (i.e., nums[j] - nums[i]), such that 0 <= i < j <
		n and nums[i] < nums[j].

	Return the maximum difference. If no such i and j exists, return -1.
*/

package solutions

func maximumDifference(nums []int) int {
	res := -1
	for i, preMin := 1, nums[0]; i < len(nums); i++ {
		if nums[i] > preMin {
			res = max(res, nums[i]-preMin)
		} else {
			preMin = nums[i]
		}
	}

	return res
}
