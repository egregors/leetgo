/*
	https://leetcode.com/problems/non-decreasing-array/

	Given an array nums with n integers, your task is to check if it could become
		non-decreasing by
	modifying at most one element.

	We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every
		i (0-based) such
	that (0 <= i <= n - 2).
*/

package solutions

func checkPossibility(nums []int) bool {
	var fl bool
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if fl {
				return false
			}
			fl = true
			if i < 2 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}
	return true
}
