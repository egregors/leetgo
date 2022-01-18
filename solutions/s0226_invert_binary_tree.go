/*
	https://leetcode.com/problems/invert-binary-tree/

	Given the root of a binary tree, invert the tree, and return its root.
*/

package solutions

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
