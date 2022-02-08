/*
	https://leetcode.com/problems/binary-tree-level-order-traversal/

	Given the root of a binary tree, return the level order traversal of its nodes' values.
	(i.e., from left to right, level by level).
*/

//nolint:revive // Das ist OK
package solutions

func bfs(q *TreeNodeQ) [][]int {
	var lineRes []int
	var res [][]int
	var nextQ TreeNodeQ

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
	return bfs(&TreeNodeQ{root})
}
