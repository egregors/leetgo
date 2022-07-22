/*
	https://leetcode.com/problems/partition-list/

	Given the head of a linked list and a value x, partition it such that all nodes less than x
	come before nodes greater than or equal to x.

	You should preserve the original relative order of the nodes in each of the two partitions.
*/

package solutions

func partition(head *ListNode, x int) *ListNode {
	bHead, aHead := new(ListNode), new(ListNode)
	b, a := bHead, aHead

	for head != nil {
		if head.Val < x {
			b.Next = head
			b = b.Next
		} else {
			a.Next = head
			a = a.Next
		}
		head = head.Next
	}

	a.Next = nil
	b.Next = aHead.Next

	return bHead.Next
}
