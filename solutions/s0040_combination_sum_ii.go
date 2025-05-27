/*
	https://leetcode.com/problems/combination-sum-ii/

	Given a collection of candidate numbers (candidates) and a target number
		(target),
	find all unique combinations in candidates where the candidate numbers sum to
		target.

	Each number in candidates may only be used once in the combination.

	Note: The solution set must not contain duplicate combinations.
*/

package solutions

import "sort"

func backtrack40(idx, remain int, curr, nums []int, acc *[][]int) {
	if remain == 0 {
		cp := make([]int, len(curr))
		copy(cp, curr)
		*acc = append(*acc, cp)
		return
	} else if remain < 0 {
		return
	}

	for i := idx; i < len(nums); i++ {
		if i > idx && nums[i] == nums[i-1] {
			continue
		}

		curr = append(curr, nums[i])
		backtrack40(i+1, remain-nums[i], curr, nums, acc)
		curr = curr[:len(curr)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	var acc [][]int
	sort.Ints(candidates)
	backtrack40(0, target, []int{}, candidates, &acc)
	return acc
}
