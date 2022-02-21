/*
	https://leetcode.com/problems/balanced-binary-tree/

	Given a binary tree, determine if it is height-balanced.

	For this problem, a height-balanced binary tree is defined as:

		a binary tree in which the left and right subtrees of every node differ in height
	by no more than 1.
*/

package solutions

func dfs0110(n *TreeNode) (isBalanced bool, height int) {
	if n == nil {
		return true, -1
	}

	isLeftBalanced, leftHeight := dfs0110(n.Left)
	if !isLeftBalanced {
		return false, 0
	}
	isRightBalanced, rightHeight := dfs0110(n.Right)
	if !isRightBalanced {
		return false, 0
	}

	return Abs(leftHeight-rightHeight) < 2, 1 + Maximum(leftHeight, rightHeight)
}

func isBalanced(root *TreeNode) bool {
	res, _ := dfs0110(root)
	return res
}
