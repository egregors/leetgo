/*
	https://leetcode.com/problems/merge-k-sorted-lists

	You are given an array of k linked-lists lists, each linked-list is sorted in
	ascending order.

	Merge all the linked-lists into one sorted linked-list and return it.
*/

package solutions

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 1 {
		return lists[0]
	}

	var res *ListNode
	for _, n := range lists {
		res = merge2Lists(res, n)
	}

	return res
}

func merge2Lists(l, r *ListNode) *ListNode {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}

	var res ListNode

	if l.Val < r.Val {
		res.Val = l.Val
		res.Next = merge2Lists(l.Next, r)
	} else {
		res.Val = r.Val
		res.Next = merge2Lists(l, r.Next)
	}

	return &res
}
