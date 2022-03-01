package solutions

func mergeSort(n *ListNode) *ListNode {
	if n == nil || n.Next == nil {
		return n
	}

	mid := n.GetMiddle()
	nextOfMid := mid.Next

	mid.Next = nil
	left := mergeSort(n)
	right := mergeSort(nextOfMid)

	sortedList := sortedMerge(left, right)
	return sortedList
}

func sortedMerge(a, b *ListNode) *ListNode {
	res := (*ListNode)(nil)
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	if a.Val <= b.Val {
		res = a
		res.Next = sortedMerge(a.Next, b)
	} else {
		res = b
		res.Next = sortedMerge(a, b.Next)
	}
	return res
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}
