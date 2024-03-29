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

// Sum returns the sum of Ints
func Sum(xs ...int) int {
	s := 0
	for _, x := range xs {
		s += x
	}
	return s
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

// FindTreeNode returns TreeNode with value v if it exist, otherwise - nil
func FindTreeNode(n *TreeNode, v int) *TreeNode {
	if n.Val == v {
		return n
	}

	if n.Left != nil {
		if res := FindTreeNode(n.Left, v); res != nil {
			return res
		}
	}

	if n.Right != nil {
		if res := FindTreeNode(n.Right, v); res != nil {
			return res
		}
	}

	return nil
}

// Contains return true if some collection `xs` contains element `y`
func Contains[T comparable](xs []T, y T) bool {
	for _, x := range xs {
		if x == y {
			return true
		}
	}
	return false
}
