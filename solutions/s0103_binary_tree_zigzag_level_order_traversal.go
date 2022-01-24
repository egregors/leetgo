/*
	https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

	Given the root of a binary tree, return the zigzag level order traversal of its nodes' values.
	(i.e., from left to right, then right to left for the next level and alternate between).
*/

package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func bfs103(q []*TreeNode, isASC bool) [][]int {
	if len(q) == 0 {
		return [][]int{{}}
	}

	var nextQ []*TreeNode
	var lvl []int
	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		if n.Left != nil {
			nextQ = append(nextQ, n.Left)
		}
		if n.Right != nil {
			nextQ = append(nextQ, n.Right)
		}

		lvl = append(lvl, n.Val)
	}

	var res [][]int
	if isASC {
		rev103(&lvl)
	}
	res = append([][]int{lvl}, bfs103(nextQ, !isASC)...)
	return res
}

func rev103(xs *[]int) {
	for i, j := 0, len(*xs)-1; i < len(*xs)>>1; i, j = i+1, j-1 {
		(*xs)[i], (*xs)[j] = (*xs)[j], (*xs)[i]
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	trav := bfs103([]*TreeNode{root}, false)
	return trav[:len(trav)-1]
}
