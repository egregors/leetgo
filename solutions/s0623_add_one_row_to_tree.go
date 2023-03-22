/*
	https://leetcode.com/problems/add-one-row-to-tree/

	Given the root of a binary tree and two integers val and depth, add a row of nodes with value val at the
	given depth depth.

	Note that the root node is at depth 1.

	The adding rule is:

		Given the integer depth, for each not null tree node cur at the depth depth - 1, create two tree
			nodes with value val as cur's left subtree root and right subtree root.
		cur's original left subtree should be the left subtree of the new left subtree root.
		cur's original right subtree should be the right subtree of the new right subtree root.
		If depth == 1 that means there is no depth depth - 1 at all, then create a tree node with
			value val as the new root of the whole original tree, and the original tree is the new root's left subtree.
*/

//nolint:revive // it's ok
package solutions

type Queue623 []*TreeNode

func (q *Queue623) IsEmpty() bool    { return len(*q) == 0 }
func (q *Queue623) Push(n *TreeNode) { *q = append(*q, n) }
func (q *Queue623) Pop() *TreeNode {
	if len(*q) > 0 {
		n := (*q)[0]
		*q = (*q)[1:]
		return n
	}
	return nil
}

func addOneRow(root *TreeNode, val, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}

	var q Queue623
	q.Push(root)
	currDepth := 1

	for !q.IsEmpty() {
		var tmp, nextQ Queue623

		for !q.IsEmpty() {
			n := q.Pop()
			tmp.Push(n)
			nextQ.Push(n.Left)
			nextQ.Push(n.Right)
		}

		if currDepth == depth-1 {
			for n := tmp.Pop(); n != nil; n = tmp.Pop() {
				n.Left = &TreeNode{Val: val, Left: nextQ.Pop(), Right: nil}
				n.Right = &TreeNode{Val: val, Left: nil, Right: nextQ.Pop()}
			}
			return root
		}

		// remove nil's
		var foo Queue623
		for _, n := range nextQ {
			if n != nil {
				foo.Push(n)
			}
		}

		currDepth++
		q = foo
	}

	return root
}
