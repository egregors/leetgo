/*
	https://leetcode.com/problems/count-complete-tree-nodes/

	Given the root of a complete binary tree, return the number of the nodes in the
		tree.

	According to Wikipedia, every level, except possibly the last, is completely
		filled in a
	complete binary tree, and all nodes in the last level are as far left as
		possible. It can
	have between 1 and 2h nodes inclusive at the last level h.

	Design an algorithm that runs in less than O(n) time complexity.
*/

//nolint:revive // it's ok
package solutions

import (
	"math"
)

type Q222 []*TreeNode

func (q Q222) IsEmpty() bool     { return len(q) == 0 }
func (q *Q222) Push(n *TreeNode) { *q = append(*q, n) }
func (q *Q222) Next() *TreeNode {
	// fixme: what if q is empty
	n := (*q)[0]
	*q = (*q)[1:]
	return n
}

func count(lvl int) int {
	var res int
	for p := 0; p <= lvl; p++ {
		res += intPow(2, p)
	}
	return res
}

func intPow(n, p int) int {
	return int(math.Pow(float64(n), float64(p)))
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	dfs := func(q Q222, lvl int) int {
		for !q.IsEmpty() {
			var nextQ Q222
			c := 0
			for !q.IsEmpty() {
				n := q.Next()
				if n.Left == nil {
					return c + count(lvl)
				}
				c++
				nextQ.Push(n.Left)

				if n.Right == nil {
					return c + count(lvl)
				}
				c++
				nextQ.Push(n.Right)
			}
			q = nextQ
			lvl++
		}
		return -1
	}
	return dfs(Q222{root}, 0)
}
