/*
	https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/description/

	Given a binary array nums, you should delete one element from it.

	Return the size of the longest non-empty subarray containing only 1's in the
		resulting array.
	Return 0 if there is no such subarray.
*/

package solutions

func longestSubarray(nums []int) int {
	var l, r, c int
	for r = 0; r < len(nums); r++ {
		if nums[r] == 0 {
			c++
		}
		if c > 1 {
			if nums[l] == 0 {
				c--
			}
			l++
		}
	}

	return r - l - 1
}
