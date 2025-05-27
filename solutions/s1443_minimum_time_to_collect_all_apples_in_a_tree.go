/*
	https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/description/

	Given an undirected tree consisting of n vertices numbered from 0 to n-1, which
		has some apples
	in their vertices. You spend 1 second to walk over one edge of the tree. Return
		the minimum time in
	seconds you have to spend to collect all apples in the tree, starting at vertex
		0 and coming back to
	this vertex.

	The edges of the undirected tree are given in the array edges, where edges[i] =
		[ai, bi] means that
	exists an edge connecting the vertices ai and bi. Additionally, there is a
		boolean array hasApple,
	where hasApple[i] = true means that vertex i has an apple; otherwise, it does
		not have any apple.
*/

package solutions

func minTime(_ int, edges [][]int, hasApple []bool) int {
	var dfs func(node, parent int, adj map[int][]int) int
	dfs = func(node, parent int, adj map[int][]int) int {
		if _, ok := adj[node]; !ok {
			return 0
		}

		totalTime, childTime := 0, 0
		for _, child := range adj[node] {
			if child == parent {
				continue
			}
			childTime = dfs(child, node, adj)
			if childTime > 0 || hasApple[child] {
				totalTime += childTime + 2
			}
		}
		return totalTime
	}

	adj := make(map[int][]int)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	return dfs(0, -1, adj)
}
