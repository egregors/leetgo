/*
	https://leetcode.com/problems/reverse-nodes-in-k-group/

	Given the head of a linked list, reverse the nodes of the list k at a time, and
		return the modified list.

	k is a positive integer and is less than or equal to the length of the linked
		list. If the number of nodes
	is not a multiple of k then left-out nodes, in the end, should remain as it is.

	You may not alter the values in the list's nodes, only nodes themselves may be
		changed.
*/

package solutions

func rev25(head *ListNode) *ListNode {
	var prev, curr *ListNode = nil, head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func revGroup(head *ListNode, k int) *ListNode {
	i := 1
	curr := head
	for i < k && curr != nil {
		curr = curr.Next
		i++
	}

	if i < k || curr == nil {
		return head
	}

	tail := curr.Next
	curr.Next = nil

	newHead := rev25(head)
	head.Next = revGroup(tail, k)

	return newHead
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	return revGroup(head, k)
}
