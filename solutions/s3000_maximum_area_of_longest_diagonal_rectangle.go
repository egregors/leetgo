/*
	https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle

	You are given a 2D 0-indexed integer array dimensions.

	For all indices i, 0 <= i < dimensions.length, dimensions[i][0] represents the
		length and dimensions[i][1] represents the width of the rectangle i.

	Return the area of the rectangle having the longest diagonal. If there are
		multiple rectangles with the longest diagonal, return the area of the
		rectangle having the maximum area.
*/

package solutions

func areaOfMaxDiagonal(dimensions [][]int) int {
	maxDiaSq := 0
	maxArea := 0
	for _, dim := range dimensions {
		l := dim[0]
		w := dim[1]
		diaSq := l*l + w*w
		area := l * w
		if diaSq > maxDiaSq {
			maxDiaSq = diaSq
			maxArea = area
		} else if diaSq == maxDiaSq {
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
