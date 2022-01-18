package solutions

// ListNode is licked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Node is just a node but with Next field
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// TreeNode just a node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
