/*
	https://leetcode.com/problems/search-in-a-binary-search-tree/

	You are given the root of a binary search tree (BST) and an integer val.

	Find the node in the BST that the node's value equals val and return the
		subtree rooted with that node.
	If such a node does not exist, return null.
*/

package solutions

func bst(n *TreeNode, val int) *TreeNode {
	if n.Val == val {
		return n
	}

	if n.Left != nil {
		if r := bst(n.Left, val); r != nil {
			return r
		}
	}

	if n.Right != nil {
		if r := bst(n.Right, val); r != nil {
			return r
		}
	}

	return nil
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	return bst(root, val)
}
