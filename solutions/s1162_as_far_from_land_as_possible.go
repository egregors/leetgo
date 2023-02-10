/*
	https://leetcode.com/problems/as-far-from-land-as-possible/

	Given an n x n grid containing only values 0 and 1, where 0 represents water and 1 represents land,
	find a water cell such that its distance to the nearest land cell is maximized, and return the distance.
	If no land or water exists in the grid, return -1.

	The distance used in this problem is the Manhattan distance: the distance between two cells
	(x0, y0) and (x1, y1) is |x0 - x1| + |y0 - y1|.
*/

package solutions

func maxDistance(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var q [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, []int{i, j})
			}
		}
	}

	if len(q) == m*n {
		return -1
	}

	ans := -1
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(q) > 0 {
		ans++
		s := len(q)
		for i := 0; i < s; i++ {
			curr := q[0]
			q = q[1:]
			for _, dir := range dirs {
				x := curr[0] + dir[0]
				y := curr[1] + dir[1]
				if x >= 0 && y >= 0 && x < m && y < n && grid[x][y] == 0 {
					grid[x][y] = 1
					q = append(q, []int{x, y})
				}
			}
		}
	}
	return ans
}
