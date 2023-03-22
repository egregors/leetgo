/*
	https://leetcode.com/problems/number-of-zero-filled-subarrays

	Given an integer array nums, return the number of subarrays filled with 0.

	A subarray is a contiguous non-empty sequence of elements within an array.
*/

package solutions

func zeroFilledSubarray(nums []int) int64 {
	var res, numSub int64
	for _, n := range nums {
		if n == 0 {
			numSub++
		} else {
			numSub = 0
		}
		res += numSub
	}
	return res
}
