/*
	https://leetcode.com/problems/deepest-leaves-sum/

	Given the root of a binary tree, return the sum of values of its deepest
		leaves.
*/

//nolint:revive // it's ok
package solutions

type Queue1302 []*TreeNode

func (q *Queue1302) IsEmpty() bool    { return len(*q) == 0 }
func (q *Queue1302) Push(x *TreeNode) { *q = append(*q, x) }
func (q *Queue1302) Pop() *TreeNode {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return x
}

func bfs1302(q Queue1302, acc *[]int) {
	if q.IsEmpty() {
		return
	}

	var nextQ Queue1302
	var currSum int

	for !q.IsEmpty() {
		n := q.Pop()
		currSum += n.Val
		if n.Left != nil {
			nextQ.Push(n.Left)
		}
		if n.Right != nil {
			nextQ.Push(n.Right)
		}
	}
	*acc = append(*acc, currSum)
	bfs1302(nextQ, acc)
}

func deepestLeavesSum(root *TreeNode) int {
	var res []int
	bfs1302(Queue1302{root}, &res)
	return res[len(res)-1]
}
