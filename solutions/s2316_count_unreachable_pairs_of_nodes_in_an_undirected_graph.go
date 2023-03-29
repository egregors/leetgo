/*
	https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/

	You are given an integer n. There is an undirected graph with n nodes, numbered from
	0 to n - 1. You are given a 2D integer array edges where edges[i] = [ai, bi] denotes
	that there exists an undirected edge connecting nodes ai and bi.

	Return the number of pairs of different nodes that are unreachable from each other.
*/

package solutions

func countPairs(n int, edges [][]int) int64 {
	adjList := make([][]int, n)
	for i := range edges {
		a, b := edges[i][0], edges[i][1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a)
	}

	var dfs func(node int, visited []bool, adjList [][]int) int
	dfs = func(node int, visited []bool, adjList [][]int) int {
		visited[node] = true

		numNodesVisited := 1
		for _, child := range adjList[node] {
			if visited[child] {
				continue
			}
			numNodesVisited += dfs(child, visited, adjList)
		}
		return numNodesVisited
	}

	visited := make([]bool, n)
	var numPairs int64
	numNodesVisited := 0
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		numNodes := dfs(i, visited, adjList)
		numPairs += int64(numNodes * numNodesVisited)
		numNodesVisited += numNodes
	}
	return numPairs
}
