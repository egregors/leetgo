/*
	https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/

	You are given an m x n matrix maze (0-indexed) with empty cells (represented as
		'.') and walls
	(represented as '+'). You are also given the entrance of the maze, where
		entrance = [entrancerow, entrancecol]
	denotes the row and column of the cell you are initially standing at.

	In one step, you can move one cell up, down, left, or right. You cannot step
		into a
	cell with a wall, and you cannot step outside the maze. Your goal is to find
		the nearest exit from the entrance.
	An exit is defined as an empty cell that is at the border of the maze. The
		entrance does not count as an exit.

Return the number of steps in the shortest path from the entrance to the nearest
	exit, or -1 if no such path exists.
*/

package solutions

type t struct {
	x, y, cnt int
}

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	ds := [5]int{0, 1, 0, -1, 0}

	ex, ey, cnt := entrance[0], entrance[1], 0
	maze[ex][ey] = '+'

	q, pop := []t{}, t{}
	q = append(q, t{ex, ey, cnt})

	for len(q) > 0 {
		pop, q = q[0], q[1:]
		for i := 0; i < len(ds)-1; i++ {
			x, y := pop.x+ds[i], pop.y+ds[i+1]

			if x < 0 || x >= m || y < 0 || y >= n {
				if [2]int{pop.x, pop.y} != [2]int{ex, ey} {
					return pop.cnt
				}
			} else if maze[x][y] != '+' {
				maze[x][y] = '+'
				q = append(q, t{x, y, pop.cnt + 1})
			}
		}
	}
	return -1
}
