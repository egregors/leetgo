/*
	https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/

	Given a directed acyclic graph, with n vertices numbered from 0 to n-1, and an array edges where edges[i] = [fromi, toi]
	represents a directed edge from node fromi to node toi.

	Find the smallest set of vertices from which all nodes in the graph are reachable. It's guaranteed that a unique solution
	exists.

	Notice that you can return the vertices in any order.
*/

package solutions

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	m := make(map[int][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, 0)
	}

	for _, e := range edges {
		from, to := e[0], e[1]
		m[to] = append(m[to], from)
	}

	var res []int
	for k, v := range m {
		if len(v) == 0 {
			res = append(res, k)
		}
	}

	return res
}
