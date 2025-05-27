/*
	https://leetcode.com/problems/intersection-of-two-linked-lists/

	Given the heads of two singly linked-lists headA and headB, return the node at
		which the two lists intersect.
	If the two linked lists have no intersection at all, return null.
*/

package solutions

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
