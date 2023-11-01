/*
	https://leetcode.com/problems/find-mode-in-binary-search-tree/description/

	Given the root of a binary search tree (BST) with duplicates, return all the mode(s)
	(i.e., the most frequently occurred element) in it.

	If the tree has more than one mode, return them in any order.

	Assume a BST is defined as follows:

		The left subtree of a node contains only nodes with keys less than or equal to the node's key.
		The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
		Both the left and right subtrees must also be binary search trees.

*/

package solutions

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) (res []int) {
	var dfs func(n *TreeNode, m map[int]int)
	dfs = func(n *TreeNode, m map[int]int) {
		if n == nil {
			return
		}
		m[n.Val]++

		dfs(n.Left, m)
		dfs(n.Right, m)
	}

	m := make(map[int]int)

	dfs(root, m)

	m2 := make(map[int][]int)
	var maxCount = math.MinInt

	for k, v := range m {
		m2[v] = append(m2[v], k)
		if v > maxCount {
			maxCount = v
		}
	}

	return m2[maxCount]
}
