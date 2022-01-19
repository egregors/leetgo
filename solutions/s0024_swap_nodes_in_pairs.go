/*
	https://leetcode.com/problems/swap-nodes-in-pairs/

	Given a linked list, swap every two adjacent nodes and return its head.
	You must solve the problem without modifying the values in the list's nodes (i.e.,
	only nodes themselves may be changed.)
*/

package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	res := &ListNode{Val: -1}
	res.Next = head
	prev := res

	for head != nil && head.Next != nil {
		fst, snd := head, head.Next

		prev.Next = snd
		fst.Next = snd.Next
		snd.Next = fst

		prev = fst
		head = fst.Next
	}

	return res.Next
}
