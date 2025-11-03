/*
	https://leetcode.com/problems/count-unguarded-cells-in-the-grid

	You are given two integers m and n representing a 0-indexed m x n grid. You are
		also given two 2D integer arrays guards and walls where guards[i] = [rowi,
		coli] and walls[j] = [rowj, colj] represent the positions of the ith guard and
		jth wall respectively.

	A guard can see every cell in the four cardinal directions (north, east, south,
		or west) starting from their position unless obstructed by a wall or another
		guard. A cell is guarded if there is at least one guard that can see it.

	Return the number of unoccupied cells that are not guarded.
*/

package solutions

//nolint:revive // it is fine
func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	a := make([][]int, m)
	for ri := range a {
		a[ri] = make([]int, n)
	}

	const WALL = 1
	const GUARD = 2
	const DANGER = 3

	for _, g := range guards {
		r, c := g[0], g[1]
		a[r][c] = GUARD
	}

	for _, w := range walls {
		r, c := w[0], w[1]
		a[r][c] = WALL
	}

	check := func(r, c int) bool {
		switch a[r][c] {
		case WALL, GUARD:
			return false
		case 0:
			a[r][c] = DANGER
		}
		return true
	}

	for _, g := range guards {
		for r, c := g[0]-1, g[1]; r >= 0 && check(r, c); r-- {
		}
		for r, c := g[0]+1, g[1]; r < m && check(r, c); r++ {
		}
		for r, c := g[0], g[1]-1; c >= 0 && check(r, c); c-- {
		}
		for r, c := g[0], g[1]+1; c < n && check(r, c); c++ {
		}
	}

	cnt := 0

	for _, r := range a {
		for _, c := range r {
			if c == 0 {
				cnt++
			}
		}
	}

	return cnt
}
