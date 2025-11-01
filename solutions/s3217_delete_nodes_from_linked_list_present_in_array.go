/*
	https://leetcode.com/problems/delete-nodes-from-linked-list-present-in-array

	You are given an array of integers nums and the head of a linked list. Return
		the head of the modified linked list after removing all nodes from the linked
		list that have a value that exists in nums.
*/

package solutions

func modifiedList(nums []int, head *ListNode) *ListNode {
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}

	for head != nil && m[head.Val] {
		head = head.Next
	}

	if head == nil {
		return nil
	}

	cur := head
	for cur.Next != nil {
		if m[cur.Next.Val] {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return head
}
