/*
	https://leetcode.com/problems/check-if-array-is-sorted-and-rotated

	Given an array nums, return true if the array was originally sorted in
		non-decreasing order, then rotated some number of positions (including zero).
		Otherwise, return false.

	There may be duplicates in the original array.

	Note: An array A rotated by x positions results in an array B of the same
		length such that B[i] == A[(i+x) % A.length] for every valid index i.
*/

package solutions

func check1752(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}

	cnt := 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			cnt++
			if cnt > 1 {
				return false
			}
		}
	}

	if nums[0] < nums[n-1] {
		cnt++
	}

	return cnt <= 1
}
