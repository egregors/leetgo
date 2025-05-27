/*
	https://leetcode.com/problems/maximum-subarray/

	Given an integer array nums, find the contiguous subarray (containing at least
		one number)
	which has the largest sum and return its sum.

	A subarray is a contiguous part of an array.
*/

package solutions

func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]
	for _, v := range nums[1:] {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if max < sum {
			max = sum
		}
	}
	return max
}
