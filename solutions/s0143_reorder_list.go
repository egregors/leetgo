/*
	https://leetcode.com/problems/reorder-list/

	You are given the head of a singly linked-list. The list can be represented as:

	L0 → L1 → … → Ln - 1 → Ln

	Reorder the list to be on the following form:

	L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

	You may not modify the values in the list's nodes. Only nodes themselves may be
		changed.
*/

package solutions

func getMid(head *ListNode) *ListNode {
	var ns []*ListNode
	for curr := head; curr != nil; curr = curr.Next {
		ns = append(ns, curr)
	}
	return ns[len(ns)>>1]
}

func rev143(head *ListNode) *ListNode {
	var curr, prev *ListNode = head, nil
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func merge143(fst, snd *ListNode) *ListNode {
	h := fst
	for snd.Next != nil {
		next := fst.Next
		fst.Next = snd
		fst = next

		next = snd.Next
		snd.Next = fst
		snd = next
	}
	return h
}

func reorderList(head *ListNode) {
	merge143(head, rev143(getMid(head)))
}
