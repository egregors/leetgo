/*
	https://leetcode.com/problems/symmetric-tree/

	Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*/

package solutions

import (
	"reflect"
)

const empty = -999

func dfsPre(n *TreeNode) []int {
	val := []int{n.Val}
	if n.Left != nil {
		val = append(val, dfsPre(n.Left)...)
	} else {
		val = append(val, empty)
	}
	if n.Right != nil {
		val = append(val, dfsPre(n.Right)...)
	} else {
		val = append(val, empty)
	}
	return val
}

func dfsPost(n *TreeNode) []int {
	var val []int
	if n.Right != nil {
		val = append(val, dfsPost(n.Right)...)
	} else {
		val = append(val, empty)
	}
	if n.Left != nil {
		val = append(val, dfsPost(n.Left)...)
	} else {
		val = append(val, empty)
	}
	val = append([]int{n.Val}, val...)
	return val
}

func isSymmetric(root *TreeNode) bool {
	return reflect.DeepEqual(dfsPre(root), dfsPost(root))
}
