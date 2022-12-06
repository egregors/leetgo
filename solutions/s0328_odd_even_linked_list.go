/*
	https://leetcode.com/problems/odd-even-linked-list/

		Given the head of a singly linked list, group all the nodes with odd indices together
	followed by the nodes with even indices, and return the reordered list.

	The first node is considered odd, and the second node is even, and so on.

	Note that the relative order inside both the even and odd groups should remain as it was
	in the input.

	You must solve the problem in O(1) extra space complexity and O(n) time complexity.
*/

package solutions

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd, even := head, head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
