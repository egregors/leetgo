/*
	https://leetcode.com/problems/remove-duplicates-from-sorted-list/

	Given the head of a sorted linked list, delete all duplicates such that each element appears
	only once. Return the linked list sorted as well.
*/

package solutions

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	chars := map[int]bool{head.Val: true}
	c := head
	for c.Next != nil {
		if chars[c.Next.Val] {
			c.Next = c.Next.Next
		} else {
			chars[c.Next.Val] = true
			c = c.Next
		}
	}

	return head
}
