/*
	https://leetcode.com/problems/find-all-k-distant-indices-in-an-array

	You are given a 0-indexed integer array nums and two integers key and k. A
		k-distant index is an index i of nums for which there exists at least one
		index j such that |i - j| <= k and nums[j] == key.

	Return a list of all k-distant indices sorted in increasing order.
*/

package solutions

import "sort"

func findKDistantIndices(nums []int, key int, k int) []int {
	res := make([]int, 0, 32)

	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			for j := i - k; j <= i+k; j++ {
				m[j] = struct{}{}
			}
		}
	}

	for k := range m {
		if k >= 0 && k < len(nums) {
			res = append(res, k)
		}
	}

	sort.Ints(res)

	return res
}
