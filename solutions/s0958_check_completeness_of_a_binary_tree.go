/*
	https://leetcode.com/problems/check-completeness-of-a-binary-tree

	Given the root of a binary tree, determine if it is a complete binary tree.

	In a complete binary tree, every level, except possibly the last, is completely
		filled, and all nodes in
	the last level are as far left as possible. It can have between 1 and 2h nodes
		inclusive at the last level h.
*/

package solutions

func isCompleteTree(root *TreeNode) bool {
	fl := false
	q := []*TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node == nil {
			fl = true
		} else {
			if fl {
				return false
			}
			q = append(q, node.Left, node.Right)
		}
	}
	return true
}
