package solutions

// LinkedListLen returns amount on nodes in linked list (ListNode)
func LinkedListLen(n *ListNode) int {
	curr, c := n, 0
	for curr != nil {
		c++
		curr = curr.Next
	}
	return c
}

// Minimum return minimum element of Int slice
func Minimum(xs ...int) int {
	min := 1<<31 - 1
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

// RemoveFromIntSlice removes k-th element of Int slice in-place
// todo: wanna move to go2 with generics
func RemoveFromIntSlice(xs *[]int, k int) {
	*xs = append((*xs)[:k], (*xs)[k+1:]...)
}
