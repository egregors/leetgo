/*
	https://leetcode.com/problems/rotate-list/

	Given the head of a linked list, rotate the list to the right by k places.
*/

package solutions

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	var ls []*ListNode
	curr := head
	i := 0

	for curr != nil {
		ls = append(ls, curr)
		curr = curr.Next
		i++
	}

	ls[len(ls)-1].Next = ls[0]

	newTail := head
	for i := 0; i < len(ls)-k%len(ls)-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil

	return newHead
}
