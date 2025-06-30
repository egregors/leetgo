/*
	https://leetcode.com/problems/longest-harmonious-subsequence/

	We define a harmonious array as an array where the difference between its
		maximum value and its minimum value is exactly 1.

	Given an integer array nums, return the length of its longest harmonious
	among all its possible subsequences.
*/

package solutions

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	res := 0

	for k, v1 := range m {
		if v2, ok := m[k+1]; ok {
			res = max(res, v1+v2)
		}
	}

	return res
}
