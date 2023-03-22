/*
	https://leetcode.com/problems/shortest-path-with-alternating-colors/

	You are given an integer n, the number of nodes in a directed graph where the nodes are
	labeled from 0 to n - 1. Each edge is red or blue in this graph, and there could be
	self-edges and parallel edges.

You are given two arrays redEdges and blueEdges where:

    redEdges[i] = [ai, bi] indicates that there is a directed red edge from node ai to node bi
	in the graph, and
    blueEdges[j] = [uj, vj] indicates that there is a directed blue edge from node uj to node
	vj in the graph.

	Return an array answer of length n, where each answer[x] is the length of the shortest path
	from node 0 to node x such that the edge colors alternate along the path, or -1 if such a
	path does not exist.
*/

package solutions

type nodeWithColor struct {
	vertex int
	color  byte
}

func shortestAlternatingPaths(n int, redEdges, blueEdges [][]int) []int {
	var adj [2][][]int
	adj[0], adj[1] = buildAdj(n, redEdges), buildAdj(n, blueEdges)

	var visit [2][]int
	visit[0], visit[1] = make([]int, n), make([]int, n)
	for i := range visit[0] {
		visit[0][i] = -1
		visit[1][i] = -1
	}
	visit[0][0], visit[1][0] = 0, 0

	queue := make([]nodeWithColor, 0, 2*n)
	queue = append(queue, nodeWithColor{vertex: 0, color: 0}, nodeWithColor{vertex: 0, color: 1})

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		nextColor, nextDistance := 1^node.color, 1+visit[node.color][node.vertex]

		edges, visiting := adj[nextColor][node.vertex], visit[nextColor]
		for _, v := range edges {
			if visiting[v] < 0 {
				visiting[v] = nextDistance
				queue = append(queue, nodeWithColor{vertex: v, color: nextColor})
			}
		}
	}

	ans := make([]int, n)
	for i := range ans {
		if v1, v2 := visit[0][i], visit[1][i]; v1 < 0 || v2 < 0 {
			ans[i] = Maximum(v1, v2)
		} else {
			ans[i] = Minimum(v1, v2)
		}
	}

	return ans
}

func buildAdj(n int, edges [][]int) [][]int {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}
	return adj
}
