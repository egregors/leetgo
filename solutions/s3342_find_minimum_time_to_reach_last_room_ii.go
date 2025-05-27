/*
	https://leetcode.com/problems/find-minimum-time-to-reach-last-room-ii/description/

	There is a dungeon with n x m rooms arranged as a grid.

	You are given a 2D array moveTime of size n x m, where moveTime[i][j]
		represents the minimum time in seconds
	when you can start moving to that room. You start from the room (0, 0) at time
		t = 0 and can move to an adjacent
	room. Moving between adjacent rooms takes one second for one move and two
		seconds for the next, alternating
	between the two.

	Return the minimum time to reach the room (n - 1, m - 1).

	Two rooms are adjacent if they share a common wall, either horizontally or
		vertically.
*/

package solutions

import "math"
import "container/heap"

func minTimeToReach3342(moveTime [][]int) int {
	rows, cols := len(moveTime), len(moveTime[0])
	seen := make([][]bool, rows)
	dist := make([][]int, rows)
	for r := range rows {
		seen[r] = make([]bool, cols)
		dist[r] = make([]int, cols)
		for c := range cols {
			dist[r][c] = math.MaxInt
		}
	}

	dist[0][0] = 0
	dirs := [][]int{
		{1, 0},
		{0, -1}, {0, 1},
		{-1, 0},
	}
	q := pq3342{point3342{0, 0, 0}}
	for q.Len() > 0 {
		p := heap.Pop(&q).(point3342)
		if seen[p.r][p.c] {
			continue
		}
		seen[p.r][p.c] = true

		for _, dir := range dirs {
			nr, nc := p.r+dir[0], p.c+dir[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				dis := max(dist[p.r][p.c], moveTime[nr][nc]) + (p.r+p.c)%2 + 1
				if dis < dist[nr][nc] {
					dist[nr][nc] = dis
					heap.Push(&q, point3342{nr, nc, dis})
				}
			}
		}

	}

	return dist[rows-1][cols-1]
}

type pq3342 []point3342

func (q *pq3342) Len() int           { return len(*q) }
func (q *pq3342) Less(i, j int) bool { return (*q)[i].d < (*q)[j].d }
func (q *pq3342) Swap(i, j int)      { (*q)[i], (*q)[j] = (*q)[j], (*q)[i] }
func (q *pq3342) Push(x any)         { *q = append(*q, x.(point3342)) }
func (q *pq3342) Pop() any {
	x := (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	return x
}

type point3342 struct {
	r, c, d int
}
