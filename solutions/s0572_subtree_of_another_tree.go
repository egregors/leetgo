/*
	https://leetcode.com/problems/subtree-of-another-tree/

	Given the roots of two binary trees root and subRoot, return true if there is a
	subtree of root with the same structure and node values of subRoot and false
	otherwise.

	A subtree of a binary tree tree is a tree that consists of a node in tree and
	all of this node's descendants. The tree tree could also be considered as a
	subtree of itself.
*/

package solutions

func isSubtree(root, subTree *TreeNode) bool {
	if root == nil || subTree == nil {
		return root == subTree
	}
	if root.Val == subTree.Val && isTreeEqual(root, subTree) {
		return true
	}
	return isSubtree(root.Left, subTree) || isSubtree(root.Right, subTree)
}

func isTreeEqual(n, m *TreeNode) bool {
	if n == nil || m == nil {
		return n == m
	}
	return n.Val == m.Val &&
		isTreeEqual(n.Left, m.Left) &&
		isTreeEqual(n.Right, m.Right)
}
