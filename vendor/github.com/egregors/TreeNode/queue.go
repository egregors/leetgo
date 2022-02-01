package treenode

// NodeQueue is a simple queue of TreeNode pointers
type NodeQueue []*TreeNode

// IsEmpty returns true if the queue is empty
func (q *NodeQueue) IsEmpty() bool { return len(*q) == 0 }

// Push adds a nodes to the end of the queue. You can push a few nodes at a time.
// They will be added like a "tail":
// ] q == [a, b, c] => q.Push(d, e, f) -> q == [a, b, c, d, e, f]
func (q *NodeQueue) Push(nodes ...*TreeNode) { *q = append(*q, nodes...) }

// Pop removes the first node from the queue and returns it.
func (q *NodeQueue) Pop() *TreeNode {
	if !q.IsEmpty() {
		n := (*q)[0]
		*q = (*q)[1:]
		return n
	}
	return nil
}
