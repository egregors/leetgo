/*
	https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/description

	You are given an integer array nums and a positive integer k.

	Return the number of subarrays where the maximum element of nums appears at least k times in that subarray.

	A subarray is a contiguous sequence of elements within an array.
*/

package solutions

import "math"

func countSubarrays2962(nums []int, k int) int64 {
	var c int64

	maxN := math.MinInt
	for _, num := range nums {
		maxN = max(maxN, num)
	}

	var l, r, cnt int
	for ; r < len(nums); r++ {
		if nums[r] == maxN {
			cnt++
		}
		for cnt == k {
			if nums[l] == maxN {
				cnt--
			}
			l++
		}
		c += int64(l)
	}

	return c
}
