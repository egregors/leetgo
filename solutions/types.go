package solutions

import (
	"fmt"

	tn "github.com/egregors/TreeNode"
)

// TreeNode is just tn.TreeNode to save back compatibility
type TreeNode = tn.TreeNode

// NewTreeNode creates a new TreeNode, ignores errors
func NewTreeNode(data string) *TreeNode {
	n, _ := tn.NewTreeNode(data)
	return n
}

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

// IntHeap Heap implementation for Ints
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds element into heap
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

// Pop removes the most small element of the heap from the heap end returns this element
func (h *IntHeap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}
