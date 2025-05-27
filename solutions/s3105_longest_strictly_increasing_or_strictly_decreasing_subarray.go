/*
	https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/description/

	You are given an array of integers nums. Return the length of the longest of
		nums which is either or .
*/

package solutions

func longestMonotonicSubarray(nums []int) int {
	inc, dec, maximum := 1, 1, 0

	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			dec = 1
			inc++
		} else if nums[i-1] > nums[i] {
			inc = 1
			dec++
		} else {
			inc = 1
			dec = 1
		}
		maximum = max(maximum, max(inc, dec))
	}

	return max(maximum, max(inc, dec))
}
