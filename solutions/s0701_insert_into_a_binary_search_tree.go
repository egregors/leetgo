/*
	https://leetcode.com/problems/insert-into-a-binary-search-tree/

	You are given the root node of a binary search tree (BST) and a value to insert
		into the tree.
	Return the root node of the BST after the insertion. It is guaranteed that the
		new value does not exist in the original BST.

	Notice that there may exist multiple valid ways for the insertion, as long as
		the tree remains a BST after insertion.
	You can return any of them.
*/

package solutions

func insert(n *TreeNode, val int) *TreeNode {
	if n == nil {
		return &TreeNode{Val: val}
	}

	if n.Val < val {
		n.Right = insert(n.Right, val)
	} else {
		n.Left = insert(n.Left, val)
	}

	return n
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	return insert(root, val)
}
