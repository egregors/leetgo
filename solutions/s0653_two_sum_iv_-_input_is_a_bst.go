/*
	https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

	Given the root of a Binary Search Tree and a target number k, return true
	if there exist two elements in the BST such that their sum is equal to the
	given target.
*/

package solutions

func dfs7(n *TreeNode, nums map[int]bool, k int) bool {
	if n == nil {
		return false
	}
	if nums[k-n.Val] {
		return true
	}
	nums[n.Val] = true
	if dfs7(n.Left, nums, k) {
		return true
	}

	if dfs7(n.Right, nums, k) {
		return true
	}
	return false
}

func findTarget(root *TreeNode, k int) bool {
	return dfs7(root, make(map[int]bool), k)
}
