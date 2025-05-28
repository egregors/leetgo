/*
	https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i

	There exist two undirected trees with n and m nodes, with distinct labels in
		ranges [0, n - 1] and [0, m - 1], respectively.

	You are given two 2D integer arrays edges1 and edges2 of lengths n - 1 and m -
		1, respectively, where edges1[i] = [ai, bi] indicates that there is an edge
		between nodes ai and bi in the first tree and edges2[i] = [ui, vi] indicates
		that there is an edge between nodes ui and vi in the second tree. You are also
		given an integer k.

	Node u is target to node v if the number of edges on the path from u to v is
		less than or equal to k. Note that a node is always target to itself.

	Return an array of n integers answer, where answer[i] is the maximum possible
		number of nodes target to node i of the first tree if you have to connect one
		node from the first tree to another node in the second tree.

	Note that queries are independent from each other. That is, for every query you
		will remove the added edge before proceeding to the next query.
*/

package solutions

func maxTargetNodes(edges1 [][]int, edges2 [][]int, k int) []int {
	var dfs func(node, parent int, children [][]int, k int) int
	dfs = func(node, parent int, children [][]int, k int) int {
		if k < 0 {
			return 0
		}
		res := 1
		for _, child := range children[node] {
			if child == parent {
				continue
			}
			res += dfs(child, node, children, k-1)
		}
		return res
	}

	build := func(edges [][]int, k int) []int {
		n := len(edges) + 1
		children := make([][]int, n)
		for _, edge := range edges {
			u, v := edge[0], edge[1]
			children[u] = append(children[u], v)
			children[v] = append(children[v], u)
		}
		res := make([]int, n)
		for i := 0; i < n; i++ {
			res[i] = dfs(i, -1, children, k)
		}
		return res
	}

	n := len(edges1) + 1
	count1 := build(edges1, k)
	count2 := build(edges2, k-1)
	maxCount2 := 0
	for _, c := range count2 {
		if c > maxCount2 {
			maxCount2 = c
		}
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = count1[i] + maxCount2
	}
	return res
}
