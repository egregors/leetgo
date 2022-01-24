/*
	https://leetcode.com/problems/path-sum-ii/

	Given the root of a binary tree and an integer targetSum, return all
	root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.

	A root-to-leaf path is a path starting from the root and ending at any leaf node.
	A leaf is a node with no children.
*/

package solutions

func dfs113(n *TreeNode, currentSum, targetSum int, path []int, res *[][]int) {
	if n.Left == nil && n.Right == nil {
		if currentSum+n.Val == targetSum {
			//nolint:gocritic // i know
			fullPath := append(path, n.Val)
			cp := make([]int, len(fullPath))
			copy(cp, fullPath)
			*res = append(*res, cp)
		}
		return
	}

	if n.Left != nil {
		dfs113(n.Left, currentSum+n.Val, targetSum, append(path, n.Val), res)
	}
	if n.Right != nil {
		dfs113(n.Right, currentSum+n.Val, targetSum, append(path, n.Val), res)
	}

}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	var answer [][]int
	dfs113(root, 0, targetSum, []int{}, &answer)
	return answer
}
