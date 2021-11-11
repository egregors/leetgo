/*
	https://leetcode.com/problems/remove-linked-list-elements/

	Given the head of a linked list and an integer val, remove all the nodes of the
	linked list that has Node.val == val, and return the new head.
*/

package solutions

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		if head.Val == val {
			return nil
		}
		return head
	}

	c := head
	for c.Next != nil {
		if c.Next.Val == val {
			c.Next = c.Next.Next
		} else {
			c = c.Next
		}
	}

	if head.Val == val {
		return head.Next
	}

	return head
}
