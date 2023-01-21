/*
	https://leetcode.com/problems/non-decreasing-subsequences/

	Given an integer array nums, return all the different possible non-decreasing
	subsequences of the given array with at least two elements. You may return the
	answer in any order.
*/

package solutions

func findSubsequences(nums []int) [][]int {
	var res [][]int
	var dfs func(int, []int)
	dfs = func(l int, cur []int) {
		if len(cur) >= 2 {
			cp := make([]int, len(cur))
			copy(cp, cur)
			res = append(res, cp)
		}
		seen := make(map[int]bool, len(nums)-l)
		for r := l; r < len(nums); r++ {
			if (l > 0 && nums[r] < nums[l-1]) || seen[nums[r]] {
				continue
			}
			seen[nums[r]] = true
			dfs(r+1, append(cur, nums[r]))
		}
	}

	dfs(0, []int{})

	return res
}
