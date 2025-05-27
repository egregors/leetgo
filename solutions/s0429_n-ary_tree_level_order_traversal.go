/*
	https://leetcode.com/problems/n-ary-tree-level-order-traversal/

	Given an n-ary tree, return the level order traversal of its nodes' values.

	Nary-Tree input serialization is represented in their level order traversal,
		each group of
	children is separated by the null value (See examples).
*/

// nolint
package solutions

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Q429 []*Node

func (q *Q429) IsEmpty() bool { return len(*q) == 0 }
func (q *Q429) Push(n *Node)  { *q = append(*q, n) }
func (q *Q429) Pop() *Node {
	if !q.IsEmpty() {
		n := (*q)[0]
		*q = (*q)[1:]
		return n
	}
	return nil
}

// levelOrder429 should call levelOrder to pass LeetCode tests
func levelOrder429(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var acc [][]int
	q := Q429{root}

	for !q.IsEmpty() {
		var (
			lvl   []int
			nextQ Q429
		)

		for !q.IsEmpty() {
			n := q.Pop()
			nextQ = append(nextQ, n.Children...)
			lvl = append(lvl, n.Val)
		}

		acc = append(acc, lvl)
		q = nextQ
	}

	return acc
}
