/*
	https://leetcode.com/problems/subsets/

	Given an integer array nums of unique elements, return all possible subsets (the power set).

	The solution set must not contain duplicate subsets. Return the solution in any order.
*/

package solutions

func backtrack78(fst, goalLen int, curr []int, acc *[][]int, nums []int) {
	if len(curr) == goalLen {
		cp := make([]int, len(curr))
		copy(cp, curr)
		*acc = append(*acc, cp)
		return
	}

	for i := fst; i < len(nums); i++ {
		curr = append(curr, nums[i])
		backtrack78(i+1, goalLen, curr, acc, nums)
		curr = curr[:len(curr)-1]
	}
}

func subsets(nums []int) [][]int {
	var res [][]int

	for i := 0; i <= len(nums); i++ {
		backtrack78(0, i, []int{}, &res, nums)
	}

	return res
}
