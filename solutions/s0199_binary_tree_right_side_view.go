/*
	https://leetcode.com/problems/binary-tree-right-side-view/

	Given the root of a binary tree, imagine yourself standing on the right side of
		it,
	return the values of the nodes you can see ordered from top to bottom.
*/

package solutions

func bfs199(q []*TreeNode) [][]int {
	var nextQ []*TreeNode
	var lvl []int
	var res [][]int

	if len(q) == 0 {
		return [][]int{{}}
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if curr.Right != nil {
			nextQ = append(nextQ, curr.Right)
		}
		if curr.Left != nil {
			nextQ = append(nextQ, curr.Left)
		}

		lvl = append(lvl, curr.Val)
	}

	res = append([][]int{lvl}, bfs199(nextQ)...)

	return res
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	lvls := bfs199([]*TreeNode{root})
	var res []int
	for _, l := range lvls {
		if len(l) != 0 {
			res = append(res, l[0])
		}
	}
	return res
}
