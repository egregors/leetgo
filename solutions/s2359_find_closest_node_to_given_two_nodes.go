/*
	https://leetcode.com/problems/find-closest-node-to-given-two-nodes/description/

	You are given a directed graph of n nodes numbered from 0 to n - 1, where each node has
	at most one outgoing edge.

	The graph is represented with a given 0-indexed array edges of size n, indicating that
	there is a directed edge from node i to node edges[i]. If there is no outgoing edge from
	i, then edges[i] == -1.

	You are also given two integers node1 and node2.

	Return the index of the node that can be reached from both node1 and node2, such that
	the maximum between the distance from node1 to that node, and from node2 to that node
	is minimized. If there are multiple answers, return the node with the smallest index,
	and if no possible answer exists, return -1.

	Note that edges may contain cycles.
*/

package solutions

import "math"

func closestMeetingNode(edges []int, node1, node2 int) int {
	bfs := func(startNode int, edges, dist []int) {
		n := len(edges)
		queue := []int{startNode}
		visit := make([]bool, n)
		dist[startNode] = 0

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			if visit[node] {
				continue
			}
			visit[node] = true
			neighbor := edges[node]
			if neighbor != -1 && !visit[neighbor] {
				dist[neighbor] = dist[node] + 1
				queue = append(queue, neighbor)

			}
		}
	}

	n := len(edges)
	dist1 := make([]int, n)
	dist2 := make([]int, n)
	for i := 0; i < n; i++ {
		dist1[i] = math.MaxInt
		dist2[i] = math.MaxInt
	}

	bfs(node1, edges, dist1)
	bfs(node2, edges, dist2)

	minDistNode := -1
	minDistTillNow := math.MaxInt
	for i := 0; i < n; i++ {
		if minDistTillNow > Maximum(dist1[i], dist2[i]) {
			minDistTillNow = Maximum(dist1[i], dist2[i])
			minDistNode = i
		}
	}
	return minDistNode
}
