/*
	https://leetcode.com/problems/combination-sum-iii/

	Find all valid combinations of k numbers that sum up to n such that the
		following conditions are true:

		Only numbers 1 through 9 are used.
		Each number is used at most once.

	Return a list of all possible valid combinations. The list must not contain the
		same combination twice,
	and the combinations may be returned in any order.
*/

package solutions

import (
	"fmt"
	"sort"
)

func getKey(xs []int) string {
	sort.Ints(xs)
	return fmt.Sprintf("%v", xs)
}

func backtrack216(curr, nums []int, target, size int, acc *[][]int, seen map[string]bool) {
	if seen[getKey(curr)] {
		return
	}

	if len(curr) == size {
		if Sum(curr...) == target {
			cp := make([]int, len(curr))
			copy(cp, curr)
			*acc = append(*acc, cp)
		}
		seen[getKey(curr)] = true
		return
	}

	for i := 0; i < len(nums); i++ {
		curr = append(curr, nums[i])
		backtrack216(curr, nums[i+1:], target, size, acc, seen)
		curr = curr[:len(curr)-1]
	}
}

func combinationSum3(k, n int) [][]int {
	var acc [][]int
	backtrack216([]int{}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, n, k, &acc, make(map[string]bool))
	return acc
}
