/*
	https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

	You are given a perfect binary tree where all leaves are on the same level, and
		every parent has two children.
	The binary tree has the following definition:

	struct Node {
	  int val;
	  Node *left;
	  Node *right;
	  Node *next;
	}

	Populate each next pointer to point to its next right node. If there is no next
		right node, the next pointer
	should be set to NULL.

	Initially, all next pointers are set to NULL.
*/

//nolint:revive // Das ist Ok
package solutions

type Queue struct {
	xs []*Node
}

func (q Queue) IsEmpty() bool {
	return len(q.xs) == 0
}

func (q *Queue) Push(n *Node) {
	q.xs = append(q.xs, n)
}

func (q *Queue) Pop() *Node {
	n := q.xs[0]
	q.xs = q.xs[1:]
	return n
}

func bfsAndLink(q Queue) {
	var line []*Node
	var nextQ Queue

	if q.IsEmpty() {
		return
	}

	for !q.IsEmpty() {
		n := q.Pop()
		if n.Left != nil {
			nextQ.Push(n.Left)
		}
		if n.Right != nil {
			nextQ.Push(n.Right)
		}
		line = append(line, n)
	}

	if len(line) != 0 {
		for i := 0; i < len(line)-1; i++ {
			line[i].Next = line[i+1]
		}
	}

	bfsAndLink(nextQ)
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	bfsAndLink(Queue{xs: []*Node{root}})
	return root
}
