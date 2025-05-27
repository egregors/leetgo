/*
	https://leetcode.com/problems/diagonal-traverse-ii/description/

	Given a 2D integer array nums, return all elements of nums in diagonal order as
		shown in the below images.
	See image in link
*/

package solutions

func findDiagonalOrder(nums [][]int) []int {
	root := &TreeNode{}
	curr := root

	for ir, r := range nums {
		tmp := curr
		for ic, c := range r {
			curr.Val = c
			if ic < len(r)-1 {
				curr.Right = &TreeNode{}
				curr = curr.Right
			}
		}
		curr = tmp
		if ir < len(nums)-1 {
			curr.Left = &TreeNode{}
			curr = curr.Left
		}
	}

	var res []int

	curr = root
	var q, nextQ []*TreeNode
	q = append(q, root)

	for len(q) > 0 {
		for _, n := range q {
			if n != nil {
				res = append(res, n.Val)
				nextQ = append(nextQ, n.Left, n.Right)
			}
		}

		q = nextQ
		nextQ = []*TreeNode{}
	}

	return res
}
