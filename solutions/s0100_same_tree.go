/*
	https://leetcode.com/problems/same-tree/description/

	Given the roots of two binary trees p and q, write a function to check if they are the same or not.

	Two binary trees are considered the same if they are structurally identical, and the nodes have the
	same value.
*/

package solutions

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
