/*
	https://leetcode.com/problems/monotonic-array/description

	An array is monotonic if it is either monotone increasing or monotone
		decreasing.

	An array nums is monotone increasing if for all i <= j, nums[i] <= nums[j]. An
		array nums is
	monotone decreasing if for all i <= j, nums[i] >= nums[j].

	Given an integer array nums, return true if the given array is monotonic, or
		false otherwise.
*/

package solutions

func isMonotonic(nums []int) bool {
	inc, dec := true, true
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			dec = false
		}
		if nums[i-1] > nums[i] {
			inc = false
		}
	}

	return inc || dec
}
