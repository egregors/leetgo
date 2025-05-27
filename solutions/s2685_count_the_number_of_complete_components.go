/*
	https://leetcode.com/problems/count-the-number-of-complete-components/description/

	You are given an integer n. There is an undirected graph with n vertices,
		numbered from 0 to n - 1.
	You are given a 2D integer array edges where edges[i] = [ai, bi] denotes that
		there exists an undirected edge
	connecting vertices ai and bi.

	Return the number of complete connected components of the graph.

	A connected component is a subgraph of a graph in which there exists a path
		between any two vertices, and
	no vertex of the subgraph shares an edge with a vertex outside of the subgraph.

	A connected component is said to be complete if there exists an edge between
		every pair of its vertices.
*/

package solutions

func countCompleteComponents(n int, edges [][]int) int {
	m := make(map[int]map[int]bool)
	for i := range n {
		m[i] = make(map[int]bool)
	}

	for _, e := range edges {
		l, r := e[0], e[1]
		m[l][r] = true
		m[r][l] = true
	}

	res := 0
	seen := make(map[int]bool)
	for k, v := range m {
		if seen[k] {
			continue
		}
		seen[k] = true

		if len(v) == 0 {
			res++
			continue
		}

		good := true
		for k2 := range v {
			seen[k2] = true
			mustContains := map[int]bool{k: true}
			for k3 := range v {
				if k3 != k2 {
					mustContains[k3] = true
					seen[k3] = true
				}
			}

			if !mEq(m[k2], mustContains) {
				good = false
				break
			}
		}
		if good {
			res++
		}

	}

	return res
}

func mEq(m1, m2 map[int]bool) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
