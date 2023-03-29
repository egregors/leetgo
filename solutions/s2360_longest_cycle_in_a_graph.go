/*
	https://leetcode.com/problems/longest-cycle-in-a-graph/

	You are given a directed graph of n nodes numbered from 0 to n - 1, where each node has at most one outgoing edge.

	The graph is represented with a given 0-indexed array edges of size n, indicating that there is a directed edge
	from node i to node edges[i]. If there is no outgoing edge from node i, then edges[i] == -1.

	Return the length of the longest cycle in the graph. If no cycle exists, return -1.

	A cycle is a path that starts and ends at the same node.
*/

package solutions

func longestCycle(edges []int) int {
	answer := -1
	var dfs func(node int, edges []int, dist map[int]int, visit []bool)
	dfs = func(node int, edges []int, dist map[int]int, visit []bool) {
		visit[node] = true
		neighbor := edges[node]

		_, ok := dist[neighbor]

		if neighbor != -1 && !visit[neighbor] {
			dist[neighbor] = dist[node] + 1
			dfs(neighbor, edges, dist, visit)
		} else if neighbor != -1 && ok {
			answer = Maximum(answer, dist[node]-dist[neighbor]+1)
		}
	}
	n := len(edges)
	visit := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visit[i] {
			dist := make(map[int]int)
			dist[i] = 1
			dfs(i, edges, dist, visit)
		}
	}
	return answer
}
