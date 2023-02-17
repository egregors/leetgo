/*
	https://leetcode.com/problems/minimum-distance-between-bst-nodes

	Given the root of a Binary Search Tree (BST), return the minimum difference between
	the values of any two different nodes in the tree.
*/

package solutions

import (
	"math"
	"sort"
)

func minDiffInBST(root *TreeNode) int {
	m := make(map[int]bool)
	m[root.Val] = true

	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Left != nil {
			m[n.Left.Val] = true
			dfs(n.Left)
		}
		if n.Right != nil {
			m[n.Right.Val] = true
			dfs(n.Right)
		}
	}

	dfs(root)
	xs := make([]int, 0, len(m))
	for k := range m {
		xs = append(xs, k)
	}
	sort.Ints(xs)

	min := math.MaxInt
	for i := 1; i < len(xs); i++ {
		if xs[i]-xs[i-1] < min {
			min = xs[i] - xs[i-1]
		}
	}

	return min
}
