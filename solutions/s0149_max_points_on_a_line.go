/*
	https://leetcode.com/problems/max-points-on-a-line/description/

	Given an array of points where points[i] = [xi, yi] represents a point on the
		X-Y plane, return the
	maximum number of points that lie on the same straight line.
*/

package solutions

import "math"

func maxPoints(points [][]int) int {
	n := len(points)
	if n == 1 {
		return 1
	}
	res := 2
	for i := 0; i < n; i++ {
		m := make(map[float64]int)
		for j := 0; j < n; j++ {
			if j != i {
				m[math.Atan2(
					float64(points[j][1]-points[i][1]),
					float64(points[j][0]-points[i][0]),
				)]++
			}
		}
		for _, v := range m {
			res = Maximum(res, v+1)
		}
	}
	return res
}
