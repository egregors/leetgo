/*
	https://leetcode.com/problems/rectangle-area/

	Given the coordinates of two rectilinear rectangles in a 2D plane, return the
	total area covered by the two rectangles.

	The first rectangle is defined by its bottom-left corner (ax1, ay1) and its
	top-right corner (ax2, ay2).

	The second rectangle is defined by its bottom-left corner (bx1, by1) and
	its top-right corner (bx2, by2).
*/

package solutions

func computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	areaOfA := (ay2 - ay1) * (ax2 - ax1)
	areaOfB := (by2 - by1) * (bx2 - bx1)

	left := Maximum(ax1, bx1)
	right := Minimum(ax2, bx2)
	xOverlap := right - left

	top := Minimum(ay2, by2)
	bottom := Maximum(ay1, by1)
	yOverlap := top - bottom

	areaOfOverlap := 0
	if xOverlap > 0 && yOverlap > 0 {
		areaOfOverlap = xOverlap * yOverlap
	}
	totalArea := areaOfA + areaOfB - areaOfOverlap
	return totalArea
}
