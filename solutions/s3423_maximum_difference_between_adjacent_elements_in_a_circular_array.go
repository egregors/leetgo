/*
	https://leetcode.com/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/

	Given a circular array nums, find the maximum absolute difference between
		adjacent elements.

	Note: In a circular array, the first and last elements are adjacent.
*/

package solutions

func maxAdjacentDistance(nums []int) int {
	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}
	diff := abs(nums[0] - nums[len(nums)-1])
	for i := 1; i < len(nums); i++ {
		diff = max(diff, abs(nums[i-1]-nums[i]))
	}

	return diff
}
