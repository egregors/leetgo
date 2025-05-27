/*
	https://leetcode.com/problems/binary-tree-preorder-traversal/

	Given the root of a binary tree, return the preorder traversal of its nodes'
		values.
*/

package solutions

func dfs1(n *TreeNode) []int {
	val := []int{n.Val}
	if n.Left != nil {
		val = append(val, dfs1(n.Left)...)
	}
	if n.Right != nil {
		val = append(val, dfs1(n.Right)...)
	}
	return val
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return dfs1(root)
}
