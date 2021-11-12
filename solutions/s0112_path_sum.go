/*
	https://leetcode.com/problems/path-sum/

	Given the root of a binary tree and an integer targetSum, return true
	if the tree has a root-to-leaf path such that adding up all the values along
	the path equals targetSum.

	A leaf is a node with no children.
*/

package solutions

func dfs5(n *TreeNode, currentSum, targetSum int) bool {
	if n.Left == nil && n.Right == nil {
		if currentSum+n.Val == targetSum {
			return true
		}
	}

	currentSum += n.Val
	if n.Left != nil {
		if dfs5(n.Left, currentSum, targetSum) {
			return true
		}
	}

	if n.Right != nil {
		if dfs5(n.Right, currentSum, targetSum) {
			return true
		}
	}

	return false
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return dfs5(root, 0, targetSum)
}
