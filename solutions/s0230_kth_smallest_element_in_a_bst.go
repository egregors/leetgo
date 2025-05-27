/*
	https://leetcode.com/problems/kth-smallest-element-in-a-bst/

	Given the root of a binary search tree, and an integer k, return the kth
		smallest
	value (1-indexed) of all the values of the nodes in the tree.
*/

package solutions

func dfs230(n *TreeNode) []int {
	var res []int

	if n.Left != nil {
		res = append(res, dfs230(n.Left)...)
	}

	res = append(res, n.Val)

	if n.Right != nil {
		res = append(res, dfs230(n.Right)...)
	}

	return res
}

func kthSmallest(root *TreeNode, k int) int {
	return dfs230(root)[k-1]
}
