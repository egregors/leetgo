/*
	https://leetcode.com/problems/find-minimum-time-to-reach-last-room-i/description/

	There is a dungeon with n x m rooms arranged as a grid.

	You are given a 2D array moveTime of size n x m, where moveTime[i][j]
		represents the minimum time in
	seconds when you can start moving to that room. You start from the room (0, 0)
		at time t = 0 and can
	move to an adjacent room. Moving between adjacent rooms takes exactly one
		second.

	Return the minimum time to reach the room (n - 1, m - 1).

	Two rooms are adjacent if they share a common wall, either horizontally or
		vertically.
*/

package solutions

import "math"
import "container/heap"

func minTimeToReach(moveTime [][]int) int {
	rows, cols := len(moveTime), len(moveTime[0])
	dist := make([][]int, rows)
	seen := make([][]bool, rows)
	for ri := range dist {
		dist[ri] = make([]int, cols)
		seen[ri] = make([]bool, cols)
		for ci := range dist[ri] {
			dist[ri][ci] = math.MaxInt
		}
	}

	dist[0][0] = 0
	q := pq3341{point3341{0, 0, 0}}
	dirs := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	for q.Len() > 0 {
		p := heap.Pop(&q).(point3341)
		if seen[p.r][p.c] {
			continue
		}
		seen[p.r][p.c] = true
		for _, dir := range dirs {
			nr, nc := p.r+dir[0], p.c+dir[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				newDist := max(dist[p.r][p.c], moveTime[nr][nc]) + 1
				if dist[nr][nc] > newDist {
					dist[nr][nc] = newDist
					heap.Push(&q, point3341{nr, nc, newDist})
				}
			}
		}
	}

	return dist[rows-1][cols-1]
}

type pq3341 []point3341

func (q *pq3341) Len() int           { return len(*q) }
func (q *pq3341) Less(i, j int) bool { return (*q)[i].d < (*q)[j].d }
func (q *pq3341) Swap(i, j int)      { (*q)[i], (*q)[j] = (*q)[j], (*q)[i] }
func (q *pq3341) Push(x any)         { *q = append(*q, x.(point3341)) }
func (q *pq3341) Pop() any {
	x := (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	return x
}

type point3341 struct {
	r, c, d int
}
