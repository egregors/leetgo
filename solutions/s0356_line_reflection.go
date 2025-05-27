/*
	https://leetcode.com/problems/line-reflection/description/

	Given n points on a 2D plane, find if there is such a line parallel to the
		y-axis that reflects the given points
	symmetrically.

	In other words, answer whether or not if there exists a line that after
		reflecting all points over the given line,
	the original points' set is the same as the reflected ones.

	Note that there can be repeated points.
*/

package solutions

import "math"

func isReflected(points [][]int) bool {
	minX, maxX := math.MaxInt, math.MinInt
	ps := make(map[[2]int]bool)

	for _, p := range points {
		ps[[2]int{p[0], p[1]}] = true
		minX = min(minX, p[0])
		maxX = max(maxX, p[0])
	}

	x := float64(minX+maxX) / float64(2)
	for _, p := range points {
		var mirror float64
		px := float64(p[0])
		if px < x {
			mirror = px + (x-px)*2
		} else if px > x {
			mirror = px - (px-x)*2
		} else {
			mirror = x
		}
		if _, ok := ps[[2]int{int(mirror), p[1]}]; !ok {
			return false
		}
	}

	return true
}
