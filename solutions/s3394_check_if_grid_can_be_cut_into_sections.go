/*
	https://leetcode.com/problems/check-if-grid-can-be-cut-into-sections/description/

	You are given an integer n representing the dimensions of an n x n grid, with
		the origin
	at the bottom-left corner of the grid. You are also given a 2D array of
		coordinates
	rectangles, where rectangles[i] is in the form [startx, starty, endx, endy],
		representing
	a rectangle on the grid. Each rectangle is defined as follows:

		(startx, starty): The bottom-left corner of the rectangle.
		(endx, endy): The top-right corner of the rectangle.

	Note that the rectangles do not overlap. Your task is to determine if it is
		possible to
	make either two horizontal or two vertical cuts on the grid such that:

		Each of the three resulting sections formed by the cuts contains at least one
			rectangle.
		Every rectangle belongs to exactly one section.

	Return true if such cuts can be made; otherwise, return false.
*/

package solutions

import "sort"

func checkValidCuts(_ int, rectangles [][]int) bool {
	return checkCuts(rectangles, 0) || checkCuts(rectangles, 1)
}

func checkCuts(rs [][]int, dim int) bool {
	c := 0
	sort.Slice(rs, func(i, j int) bool {
		return rs[i][dim] < rs[j][dim]
	})

	fEnd := rs[0][dim+2]

	for i := 1; i < len(rs); i++ {
		rect := rs[i]
		if fEnd <= rect[dim] {
			c++
		}

		fEnd = max(fEnd, rect[dim+2])
	}

	return c >= 2
}
