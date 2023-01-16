/*
	https://leetcode.com/problems/insert-interval/description/

	You are given an array of non-overlapping intervals intervals where
	intervals[i] = [starti, endi] represent the start and the end of the
	ith interval and intervals is sorted in ascending order by starti.
	You are also given an interval newInterval = [start, end] that represents
	the start and end of another interval.

	Insert newInterval into intervals such that intervals is still sorted in
	ascending order by starti and intervals still does not have any overlapping
	intervals (merge overlapping intervals if necessary).

	Return intervals after the insertion.
*/

package solutions

// insert57 should call insert to pass LeetCode tests
func insert57(intervals [][]int, newInterval []int) [][]int {
	var ans [][]int
	var i int
	var n = len(intervals)
	for i = 0; i < n && intervals[i][1] < newInterval[0]; i++ {
		ans = append(ans, intervals[i])
	}
	for ; i < n && intervals[i][0] <= newInterval[1]; i++ {
		newInterval[0] = Minimum(newInterval[0], intervals[i][0])
		newInterval[1] = Maximum(newInterval[1], intervals[i][1])
	}
	ans = append(ans, newInterval)
	for ; i < n; i++ {
		ans = append(ans, intervals[i])
	}
	return ans
}
