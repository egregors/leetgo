/*
	https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/description/

	You are given a 2D integer grid of size m x n and an integer x. In one operation, you can add x
	to or subtract x from any element in the grid.

	A uni-value grid is a grid where all the elements of it are equal.

	Return the minimum number of operations to make the grid uni-value. If it is not possible, return -1.
*/

package solutions

import "sort"

func minOperations2033(grid [][]int, x int) int {
	xs := make([]int, 0, len(grid)*len(grid[0]))
	for _, r := range grid {
		xs = append(xs, r...)
	}

	sort.Ints(xs)

	mid := xs[len(xs)/2] // 6
	c := 0
	for _, n := range xs {
		if mod(mid-n)%x != 0 {
			return -1
		}
		c += mod(mid-n) / x
	}

	return c
}

func mod(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
