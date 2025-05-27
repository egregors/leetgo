/*
	https://leetcode.com/problems/search-a-2d-matrix/

	Write an efficient algorithm that searches for a value in an m x n matrix. This
		matrix has the following properties:

    Integers in each row are sorted from left to right.
    The first integer of each row is greater than the last integer of the
    	previous row.
*/

package solutions

func bSearch(xs []int, n int) bool {
	if len(xs) < 2 {
		return xs[0] == n
	}
	if n < xs[len(xs)/2] {
		return bSearch(xs[:len(xs)/2], n)
	}
	return bSearch(xs[len(xs)/2:], n)
}

func bSearch1(xs [][]int, n int) []int {
	if len(xs) < 2 {
		return xs[0]
	}
	if n < xs[len(xs)/2][0] {
		return bSearch1(xs[:len(xs)/2], n)
	}
	return bSearch1(xs[len(xs)/2:], n)
}

func searchMatrix(matrix [][]int, target int) bool {
	return bSearch(bSearch1(matrix, target), target)
}
