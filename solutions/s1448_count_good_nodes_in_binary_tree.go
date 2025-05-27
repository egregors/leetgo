/*
	https://leetcode.com/problems/count-good-nodes-in-binary-tree/

	Given a binary tree root, a node X in the tree is named good if in the path
		from root to
	X there are no nodes with a value greater than X.

	Return the number of good nodes in the binary tree.
*/

package solutions

func goodNodes(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(n *TreeNode, max int) int {
		var good int
		if n.Val >= max {
			good = 1
		}
		newMax := Maximum(n.Val, max)
		if n.Left != nil {
			good += dfs(n.Left, newMax)
		}
		if n.Right != nil {
			good += dfs(n.Right, newMax)
		}
		return good
	}
	return dfs(root, root.Val)
}
