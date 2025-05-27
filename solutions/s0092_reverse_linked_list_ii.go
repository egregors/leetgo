/*
	Given the head of a singly linked list and two integers left and right where
		left <= right,
	reverse the nodes of the list from position left to position right, and return
		the reversed list.
*/

package solutions

func reverseBetween(head *ListNode, m, n int) *ListNode {
	dummy := new(ListNode)
	head, dummy.Next = dummy, head
	for i := 0; i < m-1; i++ {
		head = head.Next
	}
	var curr, prev *ListNode = head.Next, nil
	for i := 0; i < n-m+1; i++ {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	head.Next.Next = curr
	head.Next = prev
	return dummy.Next
}
