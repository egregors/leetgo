/*
	https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/

	Given a binary tree

	struct Node {
	  int val;
	  Node *left;
	  Node *right;
	  Node *next;
	}

	Populate each next pointer to point to its next right node. If there is no next
		right
	node, the next pointer should be set to NULL.

	Initially, all next pointers are set to NULL.
*/

package solutions

func bfs117(q []*Node) {
	if len(q) == 0 {
		return
	}

	var nextQ []*Node

	// make next q
	for _, n := range q {
		if n != nil {
			nextQ = append(nextQ, n.Left, n.Right)
		}
	}

	// connect nodes in the current q
	var curr *Node
	for i := 0; i < len(q)-1; i++ {
		if q[i] != nil {
			curr = q[i]
		}

		if curr != nil {
			curr.Next = q[i+1]
		}
	}

	bfs117(nextQ)
}

// connect117 should call connect to pass LeetCode tests
func connect117(root *Node) *Node {
	bfs117([]*Node{root})
	return root
}
