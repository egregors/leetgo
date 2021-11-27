/*
	https://leetcode.com/problems/remove-nth-node-from-end-of-list/

	Given the head of a linked list, remove the nth node from the end of the list and return its head.
*/

package solutions

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 1
	var beforeRemove *ListNode

	for node := head.Next; node != nil; node = node.Next {
		if l-n == 0 {
			beforeRemove = head
		} else if l-n > 0 {
			beforeRemove = beforeRemove.Next
		}
		l++
	}

	// first
	if beforeRemove == nil {
		return head.Next
	}

	// last
	if beforeRemove.Next.Next == nil {
		beforeRemove.Next = nil
		return head
	}

	// mid
	beforeRemove.Next = beforeRemove.Next.Next

	return head
}
