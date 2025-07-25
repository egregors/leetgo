/*
	You are given an integer array nums.

	You are allowed to delete any number of elements from nums without making it
		empty. After performing the deletions, select a

	of nums such that:

		All elements in the subarray are unique.
		The sum of the elements in the subarray is maximized.

	Return the maximum sum of such a subarray.
*/

package solutions

import "cmp"

func maxSum(nums []int) int {
	s := 0
	maxMin := -100

	m := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = struct{}{}
			s = max(s, s+n)
			if n <= 0 {
				maxMin = max(maxMin, n)
			}
		}
	}

	return cmp.Or(s, maxMin)
}
