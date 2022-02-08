/*
	https://leetcode.com/problems/shortest-path-in-binary-matrix/

	Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix.
	If there is no clear path, return -1.

	A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the
	bottom-right cell (i.e., (n - 1, n - 1)) such that:

		All the visited cells of the path are 0.
		All the adjacent cells of the path are 8-directionally connected (i.e., they are
		different and they share an edge or a corner).

	The length of a clear path is the number of visited cells of this path.
*/

//nolint:revive // meh
package solutions

type MPoint []int
type MPointQ []MPoint

func (q *MPointQ) Len() int              { return len(*q) }
func (q *MPointQ) IsEmpty() bool         { return q.Len() == 0 }
func (q *MPointQ) Push(points ...MPoint) { *q = append(*q, points...) }
func (q *MPointQ) Pop() MPoint {
	var p MPoint
	if !q.IsEmpty() {
		p = (*q)[0]
		*q = (*q)[1:]
	}
	return p
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] != 0 || grid[len(grid)-1][len(grid[0])-1] != 0 {
		return -1
	}

	var q MPointQ
	grid[0][0] = 1
	q.Push(MPoint{0, 0})

	for !q.IsEmpty() {
		p := q.Pop()
		r, c := p[0], p[1]
		dist := grid[r][c]
		if r == len(grid)-1 && c == len(grid[0])-1 {
			return dist
		}

		dirs := [][]int{
			{-1, -1}, {-1, 0}, {-1, 1},
			{0, -1}, {0, 1},
			{1, -1}, {1, 0}, {1, 1},
		}
		for _, d := range dirs {
			nextR, nextC := r+d[0], c+d[1]
			if nextR >= 0 && nextR < len(grid) && nextC >= 0 && nextC < len(grid[0]) && grid[nextR][nextC] == 0 {
				q.Push(MPoint{nextR, nextC})
				grid[nextR][nextC] = dist + 1
			}
		}
	}

	return -1
}
