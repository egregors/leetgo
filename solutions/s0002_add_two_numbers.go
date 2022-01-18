/*
	https://leetcode.com/problems/add-two-numbers/

	You are given two non-empty linked lists representing two non-negative integers.
	The digits are stored in reverse order, and each of their nodes contains a single digit.
	Add the two numbers and return the sum as a linked list.

	You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

package solutions

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head

	var carry int
	for l1 != nil || l2 != nil {
		var a, b int
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + carry
		carry = sum / 10

		curr.Next = &ListNode{
			Val: sum % 10,
		}
		curr = curr.Next
	}
	if carry > 0 {
		curr.Next = &ListNode{
			Val: carry,
		}
	}

	return head.Next
}
