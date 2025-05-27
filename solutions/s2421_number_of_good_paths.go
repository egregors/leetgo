/*
	https://leetcode.com/problems/number-of-good-paths/description/

		There is a tree (i.e. a connected, undirected graph with no cycles) consisting
			of
	n nodes numbered from 0 to n - 1 and exactly n - 1 edges.

	You are given a 0-indexed integer array vals of length n where vals[i] denotes
		the
	value of the ith node. You are also given a 2D integer array edges where
		edges[i] = [ai, bi]
	denotes that there exists an undirected edge connecting nodes ai and bi.

	A good path is a simple path that satisfies the following conditions:

		The starting node and the ending node have the same value.
		All nodes between the starting node and the ending node have values less than
			or
	equal to the starting node (i.e. the starting node's value should be the
		maximum value
	along the path).

	Return the number of distinct good paths.

	Note that a path and its reverse are counted as the same path. For example, 0
		-> 1 is
	considered to be the same as 1 -> 0. A single node is also considered as a
		valid path.
*/

package solutions

import "sort"

type unionFind struct {
	parent []int
	rank   []int
}

func (u *unionFind) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}

	return u.parent[x]
}

func (u *unionFind) Set(x, y int) {
	xSet, ySet := u.Find(x), u.Find(y)

	if xSet != ySet {
		switch {
		case u.rank[xSet] < u.rank[ySet]:
			u.parent[xSet] = ySet
		case u.rank[xSet] > u.rank[ySet]:
			u.parent[ySet] = xSet
		default:
			u.parent[ySet] = xSet
			u.rank[xSet]++
		}
	}
}

func newUnionFind(size int) *unionFind {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}

	return &unionFind{parent, make([]int, size)}
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = []int{}
	}
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}

	valuesToNodes := make(map[int][]int)
	for node := 0; node < n; node++ {
		valuesToNodes[vals[node]] = append(valuesToNodes[vals[node]], node)
	}

	keys := []int{}
	for value := range valuesToNodes {
		keys = append(keys, value)
	}

	sort.Ints(keys)

	dsu, goodPaths := newUnionFind(n), 0
	for _, value := range keys {
		nodes := valuesToNodes[value]

		for _, node := range nodes {
			for _, neighbor := range adj[node] {
				if vals[node] >= vals[neighbor] {
					dsu.Set(node, neighbor)
				}
			}
		}

		group := make(map[int]int)
		for _, node := range nodes {
			group[dsu.Find(node)]++
		}
		for _, size := range group {
			goodPaths += size * (size + 1) / 2
		}
	}

	return goodPaths
}
