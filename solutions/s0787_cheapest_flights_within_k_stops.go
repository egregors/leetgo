/*
	https://leetcode.com/problems/cheapest-flights-within-k-stops/

	There are n cities connected by some number of flights. You are given an array flights
	where flights[i] = [fromi, toi, pricei] indicates that there is a flight from city fromi
	to city toi with cost pricei.

	You are also given three integers src, dst, and k, return the cheapest price from src to
	dst with at most k stops. If there is no such route, return -1.
*/

//nolint:revive	// it's ok
package solutions

import (
	"container/heap"
	"math"
)

type Heap787 [][]int

func (h *Heap787) Len() int           { return len(*h) }
func (h *Heap787) Less(i, j int) bool { return (*h)[i][0] < (*h)[j][0] }
func (h *Heap787) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *Heap787) Push(x any)         { *h = append(*h, x.([]int)) }

func (h *Heap787) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	adj := make(map[int][][]int)
	for _, e := range flights {
		adj[e[0]] = append(adj[e[0]], []int{e[1], e[2]})
	}
	stops := make([]int, n)
	for i := 0; i < n; i++ {
		stops[i] = math.MaxInt
	}
	pq := &Heap787{}
	heap.Push(pq, []int{0, src, 0})

	for pq.Len() > 0 {
		temp := heap.Pop(pq).([]int)
		dist, node, steps := temp[0], temp[1], temp[2]
		if steps > stops[node] || steps > k+1 {
			continue
		}
		stops[node] = steps
		if node == dst {
			return dist
		}
		if _, ok := adj[node]; !ok {
			continue
		}
		for _, e := range adj[node] {
			heap.Push(pq, []int{dist + e[1], e[0], steps + 1})
		}
	}
	return -1
}
