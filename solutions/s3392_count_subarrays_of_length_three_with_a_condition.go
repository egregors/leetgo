/*
	https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/description

	Given an integer array nums, return the number of of length 3 such that the sum of the first and third numbers
	equals exactly half of the second number.
*/

package solutions

func countSubarrays3392(nums []int) int {
	c := 0
	for i := 2; i < len(nums); i++ {
		if (nums[i-2]+nums[i])*2 == nums[i-1] {
			c++
		}
	}

	return c
}
