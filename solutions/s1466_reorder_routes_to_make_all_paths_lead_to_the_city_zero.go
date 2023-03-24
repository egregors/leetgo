/*
	https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero

	There are n cities numbered from 0 to n - 1 and n - 1 roads such that there is only one way to travel between
	two different cities (this network form a tree). Last year, The ministry of transport decided to orient the
	roads in one direction because they are too narrow.

	Roads are represented by connections where connections[i] = [ai, bi] represents a road from city ai to city bi.

	This year, there will be a big event in the capital (city 0), and many people want to travel to this city.

	Your task consists of reorienting some roads such that each city can visit the city 0. Return the minimum number
	of edges changed.

	It's guaranteed that each city can reach city 0 after reorder.
*/

package solutions

func minReorder(_ int, connections [][]int) int {
	adj := make(map[int][][]int)
	for _, con := range connections {
		adj[con[0]] = append(adj[con[0]], []int{con[1], 1})
		adj[con[1]] = append(adj[con[1]], []int{con[0], 0})
	}

	var res int
	var dfs func(node, parent int, adj map[int][][]int)
	dfs = func(node, parent int, adj map[int][][]int) {
		if _, ok := adj[node]; !ok {
			return
		}
		for _, nei := range adj[node] {
			child, sign := nei[0], nei[1]
			if child != parent {
				res += sign
				dfs(child, node, adj)
			}
		}
	}

	dfs(0, -1, adj)
	return res
}
