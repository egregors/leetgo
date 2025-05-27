/*
	https://leetcode.com/problems/convert-sorted-list-to-binary-search-tree

	Given the head of a singly linked list where elements are sorted in ascending
		order, convert it to a
	height-balanced binary search tree.
*/

package solutions

func sortedListToBST(head *ListNode) *TreeNode {
	var conv func(lo, hi int) *TreeNode
	conv = func(lo, hi int) *TreeNode {
		if lo > hi {
			return nil
		}

		mid := (lo + hi) >> 1
		left := conv(lo, mid-1)

		node := &TreeNode{Val: head.Val}
		node.Left = left

		head = head.Next
		node.Right = conv(mid+1, hi)
		return node
	}

	var n int
	curr := head
	for curr != nil {
		n++
		curr = curr.Next
	}

	return conv(0, n-1)
}
