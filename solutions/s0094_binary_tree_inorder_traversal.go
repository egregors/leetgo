/*
	https://leetcode.com/problems/binary-tree-inorder-traversal/

	Given the root of a binary tree, return the inorder traversal of its nodes' values.
*/

package solutions

func dfs2(n *TreeNode) (val []int) {
	if n.Left != nil {
		val = append(val, dfs2(n.Left)...)
	}
	val = append(val, n.Val)
	if n.Right != nil {
		val = append(val, dfs2(n.Right)...)
	}
	return val
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return dfs2(root)
}
