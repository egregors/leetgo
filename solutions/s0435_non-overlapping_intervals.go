/*
	https://leetcode.com/problems/non-overlapping-intervals/

	Given an array of intervals intervals where intervals[i] = [starti, endi],
	return the minimum number of intervals you need to remove to make the rest
	of the intervals non-overlapping.a
*/

package solutions

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	removed, prev := 0, 0

	for i := 1; i < len(intervals); i++ {
		if intervals[prev][1] > intervals[i][0] {
			if intervals[prev][1] > intervals[i][1] {
				prev = i
			}
			removed++
		} else {
			prev = i
		}
	}

	return removed
}
