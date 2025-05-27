/*
	https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

	Given the root of a binary tree, flatten the tree into a "linked list":

    The "linked list" should use the same TreeNode class where the right child
    	pointer points
		to the next node in the list and the left child pointer is always null.
    The "linked list" should be in the same order as a pre-order traversal of
    	the binary tree.
*/

package solutions

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	acc := &TreeNode{Val: root.Val}
	head := acc

	var dfs func(n *TreeNode)
	dfs = func(n *TreeNode) {
		if n == nil {
			return
		}

		acc.Right = &TreeNode{Val: n.Val}
		acc = acc.Right

		if n.Left != nil {
			dfs(n.Left)
		}
		if n.Right != nil {
			dfs(n.Right)
		}
	}
	dfs(root.Left)
	dfs(root.Right)

	*root = *head
}
