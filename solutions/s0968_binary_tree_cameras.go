/*
	https://leetcode.com/problems/binary-tree-cameras/

	You are given the root of a binary tree. We install cameras on the tree nodes where
	each camera at a node can monitor its parent, itself, and its immediate children.

	Return the minimum number of cameras needed to monitor all nodes of the tree.
*/

package solutions

func dfs968(n, parent *TreeNode, cov Set[*TreeNode], ans *int) {
	if n != nil {
		dfs968(n.Left, n, cov, ans)
		dfs968(n.Right, n, cov, ans)

		if parent == nil && !cov.Contains(n) || !cov.Contains(n.Left) || !cov.Contains(n.Right) {
			*ans++
			cov.Add(n)
			cov.Add(parent)
			cov.Add(n.Left)
			cov.Add(n.Right)
		}
	}
}

func minCameraCover(root *TreeNode) int {
	cov := make(Set[*TreeNode])
	cov.Add(nil)
	var ans int
	dfs968(root, nil, cov, &ans)
	return ans
}
