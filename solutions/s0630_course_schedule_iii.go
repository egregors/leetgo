/*
	https://leetcode.com/problems/course-schedule-iii/

	There are n different online courses numbered from 1 to n.
	You are given an array courses where courses[i] = [durationi, lastDayi]
	indicate that the ith course should be taken continuously for durationi days and must be
	finished before or on lastDayi.

	You will start on the 1st day and you cannot take two or more courses simultaneously.

	Return the maximum number of courses that you can take.
*/

package solutions

import "sort"

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	var valid []int
	time := 0
	for _, c := range courses {
		if time+c[0] <= c[1] {
			valid = append(valid, c[0])
			time += c[0]
		} else {
			maxI := 0
			for i := 1; i < len(valid); i++ {
				if valid[i] > valid[maxI] {
					maxI = i
				}
			}
			if len(valid) > 0 && valid[maxI] > c[0] {
				time += c[0] - valid[maxI]
				valid[maxI] = c[0]
			}
		}
	}
	return len(valid)
}
