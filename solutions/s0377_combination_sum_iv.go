/*
	https://leetcode.com/problems/combination-sum-iv/

	Given an array of distinct integers nums and a target integer target,
	return the number of possible combinations that add up to target.

	The test cases are generated so that the answer can fit in a 32-bit
	integer.
*/

package solutions

import "sort"

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	seen := make(map[int]int)
	var bt func(remain int) int
	bt = func(remain int) int {
		if remain == 0 {
			return 1
		}
		if r, ok := seen[remain]; ok {
			return r
		}

		count := 0
		for i := 0; i < len(nums); i++ {
			r := remain - nums[i]
			if r >= 0 {
				count += bt(r)
			} else {
				break
			}
		}

		seen[remain] = count
		return count
	}

	return bt(target)
}
