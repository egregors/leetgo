/*
	https://leetcode.com/problems/maximize-the-number-of-target-nodes-after-connecting-trees-ii/

	There exist two undirected trees with n and m nodes, labeled from [0, n - 1]
		and [0, m - 1], respectively.

	You are given two 2D integer arrays edges1 and edges2 of lengths n - 1 and m -
		1, respectively, where edges1[i] = [ai, bi] indicates that there is an edge
		between nodes ai and bi in the first tree and edges2[i] = [ui, vi] indicates
		that there is an edge between nodes ui and vi in the second tree.

	Node u is target to node v if the number of edges on the path from u to v is
		even. Note that a node is always target to itself.

	Return an array of n integers answer, where answer[i] is the maximum possible
		number of nodes that are target to node i of the first tree if you had to
		connect one node from the first tree to another node in the second tree.

	Note that queries are independent from each other. That is, for every query you
		will remove the added edge before proceeding to the next query.
*/

package solutions

func maxTargetNodes3373(edges1 [][]int, edges2 [][]int) []int {
	var dfs func(node, parent, depth int, children [][]int, color []int) int
	dfs = func(node, parent, depth int, children [][]int, color []int) int {
		res := 1 - depth%2
		color[node] = depth % 2
		for _, child := range children[node] {
			if child == parent {
				continue
			}
			res += dfs(child, node, depth+1, children, color)
		}
		return res
	}

	build := func(edges [][]int, color []int) []int {
		n := len(edges) + 1
		children := make([][]int, n)
		for _, edge := range edges {
			u, v := edge[0], edge[1]
			children[u] = append(children[u], v)
			children[v] = append(children[v], u)
		}
		res := dfs(0, -1, 0, children, color)
		return []int{res, n - res}
	}

	n := len(edges1) + 1
	m := len(edges2) + 1
	color1 := make([]int, n)
	color2 := make([]int, m)
	count1 := build(edges1, color1)
	count2 := build(edges2, color2)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = count1[color1[i]] + max(count2[0], count2[1])
	}
	return res
}
