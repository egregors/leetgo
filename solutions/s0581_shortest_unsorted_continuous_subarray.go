/*
	https://leetcode.com/problems/shortest-unsorted-continuous-subarray/

	Given an integer array nums, you need to find one continuous subarray that
	if you only sort this subarray in ascending order, then the whole array will be
		sorted in ascending order.

	Return the shortest such subarray and output its length.
*/

package solutions

import "math"

func findUnsortedSubarray(nums []int) int {
	l, r := 0, len(nums)-1
	min, max := math.MaxInt, math.MinInt

	start, end := 0, -1

	for r >= 0 {
		if nums[l] >= max {
			max = nums[l]
		} else {
			end = l
		}
		if nums[r] <= min {
			min = nums[r]
		} else {
			start = r
		}
		l++
		r--
	}

	return end - start + 1
}
