/*
	https://leetcode.com/problems/merge-two-sorted-lists/

	Merge two sorted linked lists and return it as a sorted list.
	The list should be made by splicing together the nodes of the
	first two lists.
*/

//nolint:revive // it's ok
package solutions

import (
	"fmt"
)

func intSliceToList(xs []int) *ListNode {
	if len(xs) == 0 {
		return nil
	}

	r := &ListNode{}
	l := r
	i := 0

	for i < len(xs) {
		l.Val = xs[i]
		if i == len(xs)-1 {
			l.Next = nil
		} else {
			l.Next = &ListNode{}
			l = l.Next
		}
		i++
	}

	return r
}

func (l ListNode) String() string {
	next := "nil"
	if l.Next != nil {
		next = l.Next.String()
	}
	return fmt.Sprintf("%d -> %s", l.Val, next)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l2.Next, l1)
		return l2
	}
}
