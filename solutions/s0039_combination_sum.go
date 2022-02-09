/*
	https://leetcode.com/problems/combination-sum/

	Given an array of distinct integers candidates and a target integer target, return a list of all unique
	combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

	The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the
	frequency of at least one of the chosen numbers is different.

	It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for
	the given input.
*/

package solutions

func backtrack39(remain int, curr []int, fst int, acc *[][]int, candidates []int) {
	if remain == 0 {
		cp := make([]int, len(curr))
		copy(cp, curr)
		*acc = append(*acc, cp)
		return
	} else if remain < 0 {
		return
	}

	for i := fst; i < len(candidates); i++ {
		curr = append(curr, candidates[i])
		backtrack39(remain-candidates[i], curr, i, acc, candidates)
		curr = curr[:len(curr)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	backtrack39(target, []int{}, 0, &res, candidates)
	return res
}
