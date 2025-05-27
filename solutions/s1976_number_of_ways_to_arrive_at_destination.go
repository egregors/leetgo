/*
	https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/description

	You are in a city that consists of n intersections numbered from 0 to n - 1
		with
	bi-directional roads between some intersections. The inputs are generated such
	that you can reach any intersection from any other intersection and that there
	is at most one road between any two intersections.

	You are given an integer n and a 2D integer array roads where roads[i] = [ui,
		vi, timei]
	means that there is a road between intersections ui and vi that takes timei
		minutes to travel.
	You want to know in how many ways you can travel from intersection 0 to
		intersection n - 1 in
	the shortest amount of time.

	Return the number of ways you can arrive at your destination in the shortest
		amount of time.
	Since the answer may be large, return it modulo 109 + 7.
*/

package solutions

import (
	"container/heap"
	"math"
)

func countPaths(n int, roads [][]int) int {
	mod := 1_000_000_007

	graph := make([][][]int, n)
	for _, r := range roads {
		s, e, t := r[0], r[1], r[2]
		graph[s] = append(graph[s], []int{e, t})
		graph[e] = append(graph[e], []int{s, t})
	}

	h := MinHeap1976{}

	pathCount := make([]int, n)
	shortestTime := make([]int, n)
	for i := range n {
		shortestTime[i] = math.MaxInt
	}

	shortestTime[0] = 0
	pathCount[0] = 1

	heap.Init(&h)
	heap.Push(&h, []int{0, 0})
	var currTime, currNode int

	for h.Len() > 0 {
		top := heap.Pop(&h).([]int)
		currTime = top[0]
		currNode = top[1]

		if currTime > shortestTime[currNode] {
			continue
		}

		for _, n := range graph[currNode] {
			nNode, nT := n[0], n[1]
			if currTime+nT < shortestTime[nNode] {
				shortestTime[nNode] = currTime + nT
				pathCount[nNode] = pathCount[currNode]
				heap.Push(&h, []int{shortestTime[nNode], nNode})
			} else if currTime+nT == shortestTime[nNode] {
				pathCount[nNode] = (pathCount[nNode] + pathCount[currNode]) % mod
			}
		}
	}

	return pathCount[n-1]
}

type MinHeap1976 [][]int

func (h MinHeap1976) Len() int           { return len(h) }
func (h MinHeap1976) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap1976) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap1976) Push(x any)        { *h = append(*h, x.([]int)) }
func (h *MinHeap1976) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]

	return x
}
