/*
	https://leetcode.com/problems/average-of-levels-in-binary-tree/

	Given the root of a binary tree, return the average value of the nodes on each level
	in the form of an array. Answers within 10-5 of the actual answer will be accepted.
*/

//nolint:revive // it's ok
package solutions

type Q637 []*TreeNode

func (q Q637) Len() int          { return len(q) }
func (q Q637) IsEmpty() bool     { return q.Len() == 0 }
func (q *Q637) Push(n *TreeNode) { *q = append(*q, n) }

func (q *Q637) Pop() *TreeNode {
	if !q.IsEmpty() {
		n := (*q)[0]
		*q = (*q)[1:]
		return n
	}
	return nil
}

func (q Q637) Vals() []int {
	acc := make([]int, q.Len())
	for i, n := range q {
		acc[i] = n.Val
	}
	return acc
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64

	q := Q637{root}

	for !q.IsEmpty() {
		res = append(res, float64(Sum(q.Vals()...))/float64(q.Len()))

		var newQ Q637
		for n := q.Pop(); n != nil; n = q.Pop() {
			if n.Left != nil {
				newQ.Push(n.Left)
			}
			if n.Right != nil {
				newQ.Push(n.Right)
			}
		}
		q = newQ
	}

	return res
}
