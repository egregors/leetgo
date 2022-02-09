/*
	https://leetcode.com/problems/permutations-ii/

	Given a collection of numbers, nums, that might contain duplicates, return all
	possible unique permutations in any order.
*/

package solutions

import "sort"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	seen := make(map[int]bool)
	sort.Ints(nums)
	backtrack47(seen, []int{}, &res, nums)
	return res
}

func backtrack47(seen map[int]bool, curr []int, acc *[][]int, nums []int) {
	if len(curr) == len(nums) {
		cp := make([]int, len(curr))
		copy(cp, curr)
		*acc = append(*acc, cp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if seen[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !seen[i-1] {
			continue
		}
		curr = append(curr, nums[i])
		seen[i] = true
		backtrack47(seen, curr, acc, nums)
		seen[i] = false
		curr = curr[:len(curr)-1]
	}
}
