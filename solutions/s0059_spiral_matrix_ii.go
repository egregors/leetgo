/*
	https://leetcode.com/problems/spiral-matrix-ii/

	Given a positive integer n, generate an n x n matrix filled with elements from
		1 to n2
	in spiral order.
*/

package solutions

func generateMatrix(n int) [][]int {
	m := make([][]int, n)
	for r := 0; r < len(m); r++ {
		m[r] = make([]int, n)
	}

	row, col := 0, 0
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	currDir := 0

	m[0][0] = 1
	val := 2
	n2 := n * n

	for val <= n2 {
		for row+dirs[currDir][0] >= 0 && row+dirs[currDir][0] < n &&
			col+dirs[currDir][1] >= 0 && col+dirs[currDir][1] < n &&
			m[row+dirs[currDir][0]][col+dirs[currDir][1]] == 0 {
			row += dirs[currDir][0]
			col += dirs[currDir][1]
			m[row][col] = val
			val++
		}
		currDir = (currDir + 1) % 4
	}

	return m
}
