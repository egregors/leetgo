/*
	https://leetcode.com/problems/reverse-linked-list/

	Given the head of a singly linked list, reverse the list, and return the reversed list.
*/

package solutions

func rev(n *ListNode) *ListNode {
	var prev, curr, next *ListNode = nil, n, nil //nolint:ineffassign // meh
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func reverseList(head *ListNode) *ListNode {
	return rev(head)
}
