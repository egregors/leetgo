/*
	https://leetcode.com/problems/range-sum-of-bst/description/

	Given the root node of a binary search tree and two integers low and high,
		return the sum of values of all
	nodes with a value in the inclusive range [low, high].
*/

package solutions

func rangeSumBST(root *TreeNode, low int, high int) int {
	var dfs func(*TreeNode) int
	dfs = func(n *TreeNode) (sum int) {
		if n == nil {
			return 0
		}
		if low <= n.Val && n.Val <= high {
			sum += n.Val
		}
		return sum + dfs(n.Left) + dfs(n.Right)
	}
	return dfs(root)
}
