/*
	https://leetcode.com/problems/spiral-matrix/

	Given an m x n matrix, return all elements of the matrix in spiral order.
*/

package solutions

func spiralOrder(matrix [][]int) []int {
	rLen, cLen := len(matrix), len(matrix[0])
	row, col := 0, 0
	size := rLen * cLen
	visited := 101
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	currDir := 0
	res := []int{matrix[0][0]}
	matrix[0][0] = visited
	for len(res) < size {
		for row+directions[currDir][0] >= 0 && row+directions[currDir][0] < rLen &&
			col+directions[currDir][1] >= 0 && col+directions[currDir][1] < cLen &&
			matrix[row+directions[currDir][0]][col+directions[currDir][1]] != visited {
			row += directions[currDir][0]
			col += directions[currDir][1]
			res = append(res, matrix[row][col])
			matrix[row][col] = visited
			if len(res) == size {
				return res
			}
		}
		currDir = (currDir + 1) % 4
	}
	return res
}
