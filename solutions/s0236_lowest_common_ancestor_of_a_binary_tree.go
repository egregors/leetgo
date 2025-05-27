/*
	https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

	Given a binary tree, find the lowest common ancestor (LCA) of two given nodes
		in the tree.

	According to the definition of LCA on Wikipedia: “The lowest common ancestor
		is defined between
	two nodes p and q as the lowest node in T that has both p and q as descendants
		(where we allow a
	node to be a descendant of itself).”
*/

package solutions

func rec236(n, p, q, res *TreeNode) bool {
	if n == nil {
		return false
	}

	var l, r, mid int

	if rec236(n.Left, p, q, res) {
		l = 1
	}
	if rec236(n.Right, p, q, res) {
		r = 1
	}
	if (n == p) || (n == q) {
		mid = 1
	}

	if mid+l+r >= 2 {
		*res = *n
	}

	return mid+l+r > 0
}

func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	res := &TreeNode{}
	rec236(root, p, q, res)
	return res
}
