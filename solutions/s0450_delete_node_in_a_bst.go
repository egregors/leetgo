/*
	https://leetcode.com/problems/delete-node-in-a-bst/

	Given a root node reference of a BST and a key, delete the node with the given key in the BST.
	Return the root node reference (possibly updated) of the BST.

	Basically, the deletion can be divided into two stages:

		Search for a node to remove.
		If the node is found, delete the node.
*/

package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getSucc(n *TreeNode) int {
	n = n.Right
	for n.Left != nil {
		n = n.Left
	}
	return n.Val
}

func getPred(n *TreeNode) int {
	n = n.Left
	for n.Right != nil {
		n = n.Right
	}
	return n.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil {
			root.Val = getSucc(root)
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			root.Val = getPred(root)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}

	return root
}
