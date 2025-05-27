/*
	https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/

	You are given the head of a linked list. Delete the middle node, and return the
		head of the modified linked list.

	The middle node of a linked list of size n is the ⌊n / 2⌋th node from the
		start using 0-based indexing, where ⌊x⌋ denotes
		the largest integer less than or equal to x.

		For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2,
			respectively.
*/

package solutions

func deleteMiddle(head *ListNode) *ListNode {
	curr := head

	var ns []*ListNode
	for curr != nil {
		ns = append(ns, curr)
		curr = curr.Next
	}

	if len(ns) == 1 {
		return nil
	}

	if len(ns) < 3 {
		head.Next = nil
		return head
	}

	mid := len(ns) / 2
	ns[mid-1].Next = ns[mid+1]

	return head
}
