/*
	https://leetcode.com/problems/maximum-depth-of-binary-tree/

	Given the root of a binary tree, return its maximum depth.

	A binary tree's maximum depth is the number of nodes along the longest path
	from the root node down to the farthest leaf node.
*/

package solutions

func dfs4(n *TreeNode) int {
	dL, dR := 1, 1
	if n.Left != nil {
		dL += dfs4(n.Left)
	}
	if n.Right != nil {
		dR += dfs4(n.Right)
	}
	if dL > dR {
		return dL
	}
	return dR
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs4(root)
}
