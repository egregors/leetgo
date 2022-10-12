/*
	https://leetcode.com/problems/largest-perimeter-triangle/

	Given an integer array nums, return the largest perimeter of a triangle with a
	non-zero area, formed from three of these lengths. If it is impossible to form any
	triangle of a non-zero area, return 0.
*/

package solutions

import "sort"

func largestPerimeter(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	for i := len(nums) - 3; i >= 0; i-- {
		if nums[i]+nums[i+1] > nums[i+2] {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}
	return 0
}
