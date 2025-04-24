/*
	https://leetcode.com/problems/count-complete-subarrays-in-an-array/description

		You are given an array nums consisting of positive integers.

	We call a subarray of an array complete if the following condition is satisfied:

		The number of distinct elements in the subarray is equal to the number of distinct elements in the whole array.

	Return the number of complete subarrays.

	A subarray is a contiguous non-empty part of an array.
*/

package solutions

func countCompleteSubarrays(nums []int) int {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}

	goalLen := len(m)
	c := 0

	cur := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cur[nums[i]]++
		if len(cur) == goalLen {
			c++
		}
		for j := i + 1; j < len(nums); j++ {
			cur[nums[j]]++
			if len(cur) == goalLen {
				c++
			}
		}
		cur = make(map[int]int)
	}

	return c
}
