package solutions

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

// Maximum returns maximum element of Int slice
func Maximum(xs ...int) int {
	max := -1 << 31
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// LinkedListLen returns amount on nodes in linked list (ListNode)
func LinkedListLen(n *ListNode) int {
	curr, c := n, 0
	for curr != nil {
		c++
		curr = curr.Next
	}
	return c
}

// RemoveFromIntSlice removes k-th element of Int slice in-place
// todo: wanna move to go2 with generics
func RemoveFromIntSlice(xs *[]int, k int) {
	*xs = append((*xs)[:k], (*xs)[k+1:]...)
}
