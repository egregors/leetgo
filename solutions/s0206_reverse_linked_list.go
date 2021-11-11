/*
	https://leetcode.com/problems/reverse-linked-list/

	Given the head of a singly linked list, reverse the list, and return the reversed list.
*/

package solutions

func rev(n, prev *ListNode) *ListNode {
	if n == nil {
		return prev
	}
	next := n.Next
	n.Next = prev
	return rev(next, n)
}

func reverseList(head *ListNode) *ListNode {
	return rev(head, nil)
}
