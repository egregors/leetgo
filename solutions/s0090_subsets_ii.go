/*
	https://leetcode.com/problems/subsets-ii/

	Given an integer array nums that may contain duplicates, return all possible
		subsets (the power set).

	The solution set must not contain duplicate subsets. Return the solution in any
		order.
*/

package solutions

import (
	"sort"
)

func backtrack90(fst int, acc *[][]int, curr, nums []int) {
	cp := make([]int, len(curr))
	copy(cp, curr)
	*acc = append(*acc, cp)

	for i := fst; i < len(nums); i++ {
		if i != fst && nums[i] == nums[i-1] {
			continue
		}
		curr = append(curr, nums[i])
		backtrack90(i+1, acc, curr, nums)
		curr = curr[:len(curr)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var acc [][]int
	backtrack90(0, &acc, []int{}, nums)
	return acc
}
