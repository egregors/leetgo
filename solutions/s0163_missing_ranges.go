/*
	https://leetcode.com/problems/missing-ranges/

	You are given an inclusive range [lower, upper] and a sorted unique integer
		array nums, where all elements are within the inclusive range.

	A number x is considered missing if x is in the range [lower, upper] and x is
		not in nums.

	Return the shortest sorted list of ranges that exactly covers all the missing
		numbers. That is, no element of nums is included in any of the ranges, and
		each missing number is covered by one of the ranges.
*/

package solutions

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	if len(nums) == 0 {
		return [][]int{{lower, upper}}
	}
	if lower == upper {
		return [][]int{}
	}

	// [0,1,3,50,75]
	// [[2,2],[4,49],[51,74],[76,99]]
	res := make([][]int, 0)
	if lower < nums[0] {
		res = append(res, []int{lower, nums[0] - 1})
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] <= 1 {
			continue
		}
		res = append(res, []int{nums[i] + 1, nums[i+1] - 1})
	}
	if upper > nums[len(nums)-1] {
		res = append(res, []int{nums[len(nums)-1] + 1, upper})
	}

	return res
}
