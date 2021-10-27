/*
	https://leetcode.com/problems/invert-binary-tree/

	Given the root of a binary tree, invert the tree, and return its root.
*/

package solutions

// TreeNode just a node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
