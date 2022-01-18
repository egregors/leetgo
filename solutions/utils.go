package solutions

// LinkedListLen returns amount on nodes in linked list (ListNode)
func LinkedListLen(n *ListNode) int {
	curr, c := n, 0
	for curr != nil {
		c++
		curr = curr.Next
	}
	return c
}
