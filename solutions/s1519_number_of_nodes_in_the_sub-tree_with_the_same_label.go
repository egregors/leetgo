/*
	https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/

	You are given a tree (i.e. a connected, undirected graph that has no cycles)
		consisting of n
	nodes numbered from 0 to n - 1 and exactly n - 1 edges. The root of the tree is
		the node 0,
	and each node of the tree has a label which is a lower-case character given in
		the string labels
	(i.e. The node with the number i has the label labels[i]).

	The edges array is given on the form edges[i] = [ai, bi], which means there is
		an edge between nodes
	ai and bi in the tree.

	Return an array of size n where ans[i] is the number of nodes in the subtree of
		the ith node which
	have the same label as node i.

	A subtree of a tree T is the tree consisting of a node in T and all of its
		descendant nodes.
*/

package solutions

func countSubTrees(n int, edges [][]int, labels string) []int {
	adj := make(map[int][]int)
	for _, e := range edges {
		a, b := e[0], e[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	ans := make([]int, n)
	var dfs func(node, parent int) []int
	dfs = func(node, parent int) []int {
		nodeCounts := make([]int, 26)
		nodeCounts[labels[node]-'a'] = 1

		if _, ok := adj[node]; !ok {
			return nodeCounts
		}
		for _, ch := range adj[node] {
			if ch == parent {
				continue
			}
			childCounts := dfs(ch, node)
			for i := 0; i < 26; i++ {
				nodeCounts[i] += childCounts[i]
			}
		}
		ans[node] = nodeCounts[labels[node]-'a']
		return nodeCounts
	}

	dfs(0, -1)
	return ans
}
