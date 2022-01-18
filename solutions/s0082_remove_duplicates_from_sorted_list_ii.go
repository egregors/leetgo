/*
	https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

	Given the head of a sorted linked list, delete all nodes that have duplicate numbers,
	leaving only distinct numbers from the original list. Return the linked list sorted as well.
*/

package solutions

func deleteDuplicatesII(head *ListNode) *ListNode {
	curr := head
	m := make(map[int][]*ListNode)

	for curr != nil {
		m[curr.Val] = append(m[curr.Val], curr)
		curr = curr.Next
	}

	toRemove := make(map[*ListNode]bool)
	for _, v := range m {
		if len(v) > 1 {
			for _, n := range v {
				toRemove[n] = true
			}
		}
	}

	curr = head
	var prev *ListNode

	for curr != nil {
		if toRemove[curr] {
			if prev != nil && curr.Next != nil { // mid
				prev.Next = curr.Next
				curr = prev
			} else if prev == nil { // head
				if head.Next == nil {
					head = nil
					break
				}
				head = head.Next
				curr = head
				continue
			} else { // last
				prev.Next = nil
				break
			}
		}
		prev = curr
		curr = curr.Next
	}

	return head
}
