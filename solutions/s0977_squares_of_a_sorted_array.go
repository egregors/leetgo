/*
	https://leetcode.com/problems/squares-of-a-sorted-array/

	Given an integer array nums sorted in non-decreasing order,
	return an array of the squares of each number sorted in non-decreasing order.
*/

//nolint:revive //it's ok
package solutions

func sortedSquares(nums []int) []int {
	lo, hi, idx := 0, len(nums)-1, len(nums)-1
	res := make([]int, len(nums))
	for lo <= hi {
		if nums[lo]*-1 > nums[hi] {
			res[idx] = nums[lo] * nums[lo]
			lo++
		} else {
			res[idx] = nums[hi] * nums[hi]
			hi--
		}
		idx--
	}
	return res
}
