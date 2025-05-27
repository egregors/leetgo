/*
	https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/

	You are given an array of unique integers salary where salary[i] is the salary
		of the ith employee.

	Return the average salary of employees excluding the minimum and maximum
		salary. Answers within 10-5 of the actual answer will be accepted.
*/

package solutions

import "math"

func average(salary []int) float64 {
	n := len(salary)
	min, max := math.MaxInt, math.MinInt
	sum := 0

	for _, s := range salary {
		if s < min {
			min = s
		}
		if s > max {
			max = s
		}
		sum += s
	}

	return float64(sum-min-max) / float64(n-2)
}
