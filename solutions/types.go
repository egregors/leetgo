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

// TreeNodeQ is FIFO queue for TreeNode's
type TreeNodeQ []*TreeNode

// Len return amount of nodes in the queue
func (q *TreeNodeQ) Len() int { return len(*q) }

// IsEmpty return false if queue has no nodes
func (q *TreeNodeQ) IsEmpty() bool { return q.Len() == 0 }

// Push adds a node into the queue
func (q *TreeNodeQ) Push(nodes ...*TreeNode) { *q = append(*q, nodes...) }

// Pop dequeued a node from the queue
func (q *TreeNodeQ) Pop() *TreeNode {
	var n *TreeNode
	if !q.IsEmpty() {
		n = (*q)[0]
		*q = (*q)[1:]
	}
	return n
}

// ListNode is licked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode builds linked list of Ints from Ints slice
func NewListNode(xs []int) *ListNode {
	if len(xs) == 0 {
		return nil
	}

	head := &ListNode{xs[0], nil}
	curr := head
	for _, x := range xs[1:] {
		n := &ListNode{
			x,
			nil,
		}
		curr.Next = n
		curr = n
	}
	return head
}

func (n *ListNode) String() string {
	next := "nil"
	if n.Next != nil {
		next = n.Next.String()
	}
	return fmt.Sprintf("%d -> %s", n.Val, next)
}

// GetMiddle returns a middle node of linked list
func (n *ListNode) GetMiddle() *ListNode {
	slow, fast := n, n
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// ToSlice accumulate linked list values and append them into int slice
func (n *ListNode) ToSlice() []int {
	var vals []int
	curr := n
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

// Set represents a set data structure above of hashmap
type Set[K comparable] map[K]struct{}

// Add adds element into the Set
func (s Set[K]) Add(k K) { s[k] = struct{}{} }

// Remove deletes element form the Set
func (s Set[K]) Remove(k K) { delete(s, k) }

// Contains returns true if Set contains element x
func (s Set[K]) Contains(k K) bool { _, ok := s[k]; return ok }
