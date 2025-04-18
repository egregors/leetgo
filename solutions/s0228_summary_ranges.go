/*
	https://leetcode.com/problems/summary-ranges/description/

	You are given a sorted unique integer array nums.

	A range [a,b] is the set of all integers from a to b (inclusive).

	Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element
	of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges
	but not in nums.

	Each range [a,b] in the list should be output as:

		"a->b" if a != b
		"a" if a == b
*/

package solutions

import "fmt"

func summaryRanges(nums []int) []string {
	res := []string{}

	for i := 0; i < len(nums); i++ {
		start := nums[i]
		for i+1 < len(nums) && nums[i]+1 == nums[i+1] {
			i++
		}

		if start != nums[i] {
			res = append(res, fmt.Sprintf("%d->%d", start, nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%d", start))
		}

	}

	return res
}
