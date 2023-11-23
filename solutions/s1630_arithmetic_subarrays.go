/*
	https://leetcode.com/problems/arithmetic-subarrays/description/

	A sequence of numbers is called arithmetic if it consists of at least two elements,
	and the difference between every two consecutive elements is the same. More formally,
	a sequence s is arithmetic if and only if s[i+1] - s[i] == s[1] - s[0] for all valid i.

	For example, these are arithmetic sequences:

	1, 3, 5, 7, 9
	7, 7, 7, 7
	3, -1, -5, -9

	The following sequence is not arithmetic:

	1, 1, 2, 5, 7

	You are given an array of n integers, nums, and two arrays of m integers each, l and r, representing the m
	range queries, where the ith query is the range [l[i], r[i]]. All the arrays are 0-indexed.

	Return a list of boolean elements answer, where answer[i] is true if the subarray nums[l[i]],
	nums[l[i]+1], ... , nums[r[i]] can be rearranged to form an arithmetic sequence, and false otherwise.
*/

package solutions

import "sort"

func checkArithmeticSubarrays(nums, l, r []int) []bool {
	res := make([]bool, len(l))
	for i := 0; i < len(res); i++ {
		res[i] = true
	}

	for i := 0; i < len(l); i++ {
		sub := make([]int, r[i]+1-l[i])
		copy(sub, nums[l[i]:r[i]+1])
		sort.Ints(sub)

		gap := sub[1] - sub[0]
		for j := 2; j < len(sub); j++ {
			if sub[j]-sub[j-1] != gap {
				res[i] = false
				break
			}
		}
	}

	return res
}
