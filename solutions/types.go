package solutions

import "fmt"

// ListNode is licked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	next := "nil"
	if l.Next != nil {
		next = l.Next.String()
	}
	return fmt.Sprintf("%d -> %s", l.Val, next)
}

// ToSlice accumulate linked list values and append them into int slice
func (l *ListNode) ToSlice() []int {
	var vals []int
	curr := l
	for curr != nil {
		vals = append(vals, curr.Val)
		curr = curr.Next
	}
	return vals
}

// Node is just a node but with Next field
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// TreeNode just a node of binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t TreeNode) String() string {
	return fmt.Sprintf("%d %s %s", t.Val, t.Left, t.Right)
}
