/*
	https://leetcode.com/problems/move-zeroes/

	Given an integer array nums, move all 0's to the end of it while maintaining
	the relative order of the non-zero elements.

	Note that you must do this in-place without making a copy of the array.
*/

package solutions

func moveZeroes(nums []int) {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[n] = nums[n], nums[i]
			n++
		}
	}
}
