/*
	https://leetcode.com/problems/binary-tree-pruning/

	Given the root of a binary tree, return the same tree where every subtree
	(of the given tree) not containing a 1 has been removed.

	A subtree of a node node is node plus every node that is a descendant of node.
*/

package solutions

func pruneTree(root *TreeNode) *TreeNode {
	if hasOne(root) {
		return root
	}
	return nil
}

func hasOne(n *TreeNode) bool {
	if n == nil {
		return false
	}
	leftHasOne := hasOne(n.Left)
	rightHasOne := hasOne(n.Right)
	if !leftHasOne {
		n.Left = nil
	}
	if !rightHasOne {
		n.Right = nil
	}
	return n.Val == 1 || leftHasOne || rightHasOne
}
