/*
	https://leetcode.com/problems/binary-tree-level-order-traversal/

	Given the root of a binary tree, return the level order traversal of its nodes' values.
	(i.e., from left to right, level by level).
*/

//nolint:revive // Das ist OK
package solutions

type Q struct {
	xs []*TreeNode
}

func (q *Q) Push(n *TreeNode) {
	q.xs = append(q.xs, n)
}

func (q *Q) Pop() *TreeNode {
	el := q.xs[0]
	q.xs = q.xs[1:]
	return el
}

func (q Q) IsEmpty() bool {
	return len(q.xs) == 0
}

func bfs(q *Q) [][]int {
	var lineRes []int
	var res [][]int
	var nextQ Q

	if q == nil {
		return [][]int{{}}
	}

	for !q.IsEmpty() {
		n := q.Pop()
		if n.Left != nil {
			nextQ.Push(n.Left)
		}
		if n.Right != nil {
			nextQ.Push(n.Right)
		}
		lineRes = append(lineRes, n.Val)
	}
	if len(lineRes) != 0 {
		res = append([][]int{lineRes}, bfs(&nextQ)...)
	}
	return res
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	return bfs(&Q{xs: []*TreeNode{root}})
}
