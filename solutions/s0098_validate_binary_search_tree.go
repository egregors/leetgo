/*
	https://leetcode.com/problems/validate-binary-search-tree/

	Given the root of a binary tree, determine if it is a valid binary search tree
		(BST).

	A valid BST is defined as follows:

		The left subtree of a node contains only nodes with keys less than the node's
			key.
		The right subtree of a node contains only nodes with keys greater than the
			node's key.
		Both the left and right subtrees must also be binary search trees.
*/

package solutions

func dfs6(n, l, r *TreeNode) bool {
	if n == nil {
		return true
	}
	if l != nil && n.Val <= l.Val {
		return false
	}

	if r != nil && n.Val >= r.Val {
		return false
	}
	return dfs6(n.Left, l, n) && dfs6(n.Right, n, r)
}

func isValidBST(root *TreeNode) bool {
	return dfs6(root, nil, nil)
}
