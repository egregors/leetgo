/*
	https://leetcode.com/problems/find-duplicate-subtrees/

	Given the root of a binary tree, return all duplicate subtrees.

	For each kind of duplicate subtrees, you only need to return the root node of
		any one of them.

	Two trees are duplicate if they have the same structure with the same node
		values.
*/

package solutions

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	h := make(map[string]bool)
	var duplicates []*TreeNode
	var findDuplicateChild func(root *TreeNode) []byte
	findDuplicateChild = func(root *TreeNode) []byte {
		if root == nil {
			return []byte{0, 0}
		}
		val := root.Val + 201
		result := []byte{byte(val >> 8), byte(val)}
		for _, child := range [...]*TreeNode{root.Left, root.Right} {
			childResult := findDuplicateChild(child)
			if child != nil {
				if added, exist := h[string(childResult)]; exist {
					if !added {
						duplicates = append(duplicates, child)
						h[string(childResult)] = true
					}
				} else {
					h[string(childResult)] = false
				}
			}
			result = append(result, childResult...)
		}
		return result
	}
	findDuplicateChild(root)
	return duplicates
}
