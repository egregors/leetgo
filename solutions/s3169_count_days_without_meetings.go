/*
	https://leetcode.com/problems/count-days-without-meetings/description/

	You are given a positive integer days representing the total number of days an employee is available for work
	(starting from day 1). You are also given a 2D array meetings of size n where, meetings[i] = [start_i, end_i]
	represents the starting and ending days of meeting i (inclusive).

	Return the count of days when the employee is available for work but no meetings are scheduled.

	Note: The meetings may overlap.
*/

package solutions

import "sort"

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	var c, lastR int
	for _, m := range meetings {
		l, r := m[0], m[1]
		if l > lastR+1 {
			c += l - lastR - 1
		}

		lastR = max(lastR, r)
	}

	c += days - lastR
	return c
}
