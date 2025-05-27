/*
	https://leetcode.com/problems/longest-path-with-different-adjacent-characters/

	You are given a tree (i.e. a connected, undirected graph that has no cycles)
		rooted at
	node 0 consisting of n nodes numbered from 0 to n - 1. The tree is represented
		by a
	0-indexed array parent of size n, where parent[i] is the parent of node i.
		Since node
	0 is the root, parent[0] == -1.

	You are also given a string s of length n, where s[i] is the character assigned
		to node i.

	Return the length of the longest path in the tree such that no pair of adjacent
		nodes on the
	path have the same character assigned to them.
*/

package solutions

func longestPath(parent []int, s string) int {
	n := len(parent)
	res := 1

	children := make(map[int][]int)
	for i := 1; i < n; i++ {
		children[parent[i]] = append(children[parent[i]], i)
	}

	var dfs func(currNode int, s string) int
	dfs = func(currNode int, s string) int {
		if _, ok := children[currNode]; !ok {
			return 1
		}
		longestChain, secondLongestChain := 0, 0
		for _, child := range children[currNode] {
			longestChainStartingFromChild := dfs(child, s)
			if s[currNode] == s[child] {
				continue
			}
			if longestChainStartingFromChild > longestChain {
				secondLongestChain = longestChain
				longestChain = longestChainStartingFromChild
			} else if longestChainStartingFromChild > secondLongestChain {
				secondLongestChain = longestChainStartingFromChild
			}
		}
		res = Maximum(res, longestChain+secondLongestChain+1)
		return longestChain + 1
	}

	dfs(0, s)

	return res
}
