/*
	https://leetcode.com/problems/count-nodes-equal-to-average-of-subtree/

	Given the root of a binary tree, return the number of nodes where the value of
	the node is equal to the average of the values in its subtree.

	Note:

		The average of n elements is the sum of the n elements divided by n and
			rounded down to the nearest integer.
		A subtree of root is a tree consisting of root and all of its descendants.

*/

package solutions

func averageOfSubtree(root *TreeNode) int {
	var res int

	var dfs func(n *TreeNode) (int, int)
	dfs = func(n *TreeNode) (int, int) {
		if n == nil {
			return 0, 0
		}

		c1, sum1 := dfs(n.Left)
		c2, sum2 := dfs(n.Right)

		sum := sum1 + sum2 + n.Val
		c := c1 + c2 + 1

		if sum/c == n.Val {
			res++
		}

		return c, sum
	}

	dfs(root)

	return res
}
