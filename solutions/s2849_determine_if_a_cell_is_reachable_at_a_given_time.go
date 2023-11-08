/*
	https://leetcode.com/problems/determine-if-a-cell-is-reachable-at-a-given-time/description/

	You are given four integers sx, sy, fx, fy, and a non-negative integer t.

	In an infinite 2D grid, you start at the cell (sx, sy). Each second, you must move to any of its adjacent cells.

	Return true if you can reach cell (fx, fy) after exactly t seconds, or false otherwise.

	A cell's adjacent cells are the 8 cells around it that share at least one corner with it.
	You can visit the same cell several times.

*/

package solutions

func isReachableAtTime(sx, sy, fx, fy, t int) bool {
	if sx == fx && sy == fy && t == 1 {
		return false
	}

	maxD := max(Abs(fx-sx), Abs(fy-sy))

	return maxD <= t
}
