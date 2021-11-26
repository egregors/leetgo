/*
	https://leetcode.com/problems/middle-of-the-linked-list/

	Given the head of a singly linked list, return the middle node of the linked list.

	If there are two middle nodes, return the second middle node.
*/

package solutions

func middleNode(head *ListNode) *ListNode {
	l := 0
	mid := head

	for n := head; n != nil; n = n.Next {
		l++
		if l%2 == 0 {
			mid = mid.Next
		}
	}
	return mid
}
