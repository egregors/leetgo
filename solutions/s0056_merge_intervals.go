/*
	https://leetcode.com/problems/merge-intervals/

	Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
	and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/

package solutions

import (
	"sort"
)

func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{intervals[0]}
	i, i2 := 1, 0
	for i < len(intervals) {
		n1, n2 := res[i2], intervals[i]

		if n1[1] < n2[0] {
			res = append(res, n2)
			i2++
			i++
			continue
		}

		if n1[1] >= n2[0] && n1[1] < n2[1] {
			res[i2] = []int{n1[0], n2[1]}
		}

		i++
	}

	return res
}
