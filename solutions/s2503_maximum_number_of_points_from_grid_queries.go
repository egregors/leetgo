package solutions

/*
	https://leetcode.com/problems/maximum-number-of-points-from-grid-queries/description/

	You are given an m x n integer matrix grid and an array queries of size k.

	Find an array answer of size k such that for each integer queries[i] you start in the
	top left cell of the matrix and repeat the following process:

		If queries[i] is strictly greater than the value of the current cell that you are in,

	then you get one point if it is your first time visiting this cell, and you can move to
	any adjacent cell in all 4 directions: up, down, left, and right.

		Otherwise, you do not get any points, and you end this process.

	After the process, answer[i] is the maximum number of points you can get. Note that
	for each query you are allowed to visit the same cell multiple times.

	Return the resulting array answer.
*/

import (
	"container/heap"
	"sort"
)

func maxPoints2503(grid [][]int, queries []int) []int {
	res := make([]int, len(queries))
	orderedQ := make([][]int, 0, len(queries)) // [val, id]
	for i, q := range queries {
		orderedQ = append(orderedQ, []int{q, i})
	}
	sort.Slice(orderedQ, func(i, j int) bool {
		return orderedQ[i][0] < orderedQ[j][0]
	})

	h := &MinHeap2503{} // {val, r, c}
	total := 0
	seen := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		seen[i] = make([]bool, len(grid[0]))
	}
	heap.Push(h, []int{grid[0][0], 0, 0})
	seen[0][0] = true
	dirs := [][]int{{-1, 0}, {0, 1}, {0, -1}, {1, 0}}

	for _, q := range orderedQ {
		val, idx := q[0], q[1]

		for h.Len() > 0 && h.Peek()[0] < val {
			top := heap.Pop(h).([]int)
			_, cR, cC := top[0], top[1], top[2]
			total++
			for _, d := range dirs {
				newR, newC := cR+d[0], cC+d[1]
				if newR >= 0 && newR < len(grid) &&
					newC >= 0 && newC < len(grid[0]) &&
					!seen[newR][newC] {
					heap.Push(h, []int{grid[newR][newC], newR, newC})
					seen[newR][newC] = true
				}
			}
		}
		res[idx] = total
	}

	return res
}

type MinHeap2503 [][]int

func (h MinHeap2503) Len() int           { return len(h) }
func (h MinHeap2503) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap2503) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap2503) Push(x any)        { *h = append(*h, x.([]int)) }
func (h *MinHeap2503) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
func (h *MinHeap2503) Peek() []int {
	xs := heap.Pop(h).([]int)
	heap.Push(h, xs)
	return xs
}
