/*
	https://leetcode.com/problems/longest-nice-subarray/description/

	You are given an array nums consisting of positive integers.

	We call a subarray of nums nice if the bitwise AND of every pair of
	elements that are in different positions in the subarray is equal to 0.

	Return the length of the longest nice subarray.

	A subarray is a contiguous part of an array.

	Note that subarrays of length 1 are always considered nice.
*/

package solutions

func longestNiceSubarray(nums []int) int {
	var maxL, used, l, r int

	for r = 0; r < len(nums); r++ {
		for used&nums[r] != 0 {
			used ^= nums[l]
			l++
		}

		used |= nums[r]
		maxL = max(maxL, r-l+1)
	}

	return maxL
}
