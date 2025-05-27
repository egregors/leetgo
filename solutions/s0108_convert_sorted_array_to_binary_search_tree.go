/*
	https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

	Given an integer array nums where the elements are sorted in ascending order,
		convert it to a
	height-balanced binary search tree.

	A height-balanced binary tree is a binary tree in which the depth of the two
		subtrees of every
	node never differs by more than one.
*/

package solutions

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) >> 1
	n := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
	return n
}
