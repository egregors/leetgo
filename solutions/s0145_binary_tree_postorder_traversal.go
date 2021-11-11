/*
	https://leetcode.com/problems/binary-tree-postorder-traversal/

	Given the root of a binary tree, return the postorder traversal of its nodes' values.
*/

package solutions

func dfs3(n *TreeNode) []int {
	// l, r, root
	var val []int
	if n.Left != nil {
		val = append(val, dfs3(n.Left)...)
	}
	if n.Right != nil {
		val = append(val, dfs3(n.Right)...)
	}
	val = append(val, n.Val)
	return val
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return dfs3(root)
}
